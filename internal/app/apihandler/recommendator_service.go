package apihandler

import (
	"fmt"
	"sort"
	"time"

	"github.com/ducktyst/bar_recomend/internal/app/apihandler/generated/specmodels"
	"github.com/ducktyst/bar_recomend/internal/app/apihandler/generated/specops"
	"github.com/ducktyst/bar_recomend/internal/barcode/analyzer/common"
	leven "github.com/ducktyst/bar_recomend/internal/barcode/levenstein"
	"github.com/go-openapi/runtime/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type RecommendatorService struct {
	db *sqlx.DB
}

func NewRecommendatorService(db *sqlx.DB) *RecommendatorService {
	return &RecommendatorService{
		db: db,
	}
}

func (srv *RecommendatorService) GetRecommendationsBarcodeHandler(params specops.GetRecommendationsBarcodeParams) middleware.Responder {
	fmt.Println("GetRecommendationsBarcodeHandler", params.Barcode)

	res, err := srv.findByBarcode(params.Barcode)
	if err != nil {
		return specops.NewGetRecommendationsBarcodeBadRequest().WithPayload(&specmodels.GenericError{Msg: err.Error()})
	}

	payload := make([]*specmodels.Recommendation, len(res))
	for i := range payload {
		payload[i] = &specmodels.Recommendation{
			Articul:  res[i].Name,
			ShopName: res[i].ShopName,
			Barcode:  params.Barcode,
			Price:    float64(res[i].Price),
			URL:      res[i].Url,
		}
	}

	return specops.NewGetRecommendationsBarcodeOK().WithPayload(payload)
}

func (srv *RecommendatorService) PostRecommendationsHandler(params specops.PostRecommendationsParams) middleware.Responder {
	logrus.Info(time.Now().Format(time.RFC3339), " PostRecommendationsHandler start")

	if params.Content == nil { // possible?
		return specops.NewPostRecommendationsBadRequest().WithPayload(&specmodels.GenericError{Msg: "file is empty"})
	}

	img_barcode, err := common.ScanBarCodeFile(params.Content)
	if err != nil {
		return specops.NewPostRecommendationsBadRequest().WithPayload(&specmodels.GenericError{Msg: err.Error()})
	}

	logrus.Info(time.Now().Format(time.RFC3339), " PostRecommendationsHandler ", img_barcode, err)
	res, err := srv.findByBarcode(img_barcode)
	if err != nil {
		return specops.NewPostRecommendationsAnalyzeBadRequest().WithPayload(&specmodels.GenericError{Msg: err.Error()})
	}

	payload := make([]*specmodels.Recommendation, len(res))
	for i := range payload {
		payload[i] = &specmodels.Recommendation{
			Articul:  res[i].Name,
			ShopName: res[i].ShopName,
			Barcode:  img_barcode,
			Price:    float64(res[i].Price),
			URL:      res[i].Url,
		}
	}
	return specops.NewPostRecommendationsOK().WithPayload(payload)
}

func (srv *RecommendatorService) PostRecommendationsAnalyzeHandler(params specops.PostRecommendationsAnalyzeParams) middleware.Responder {
	logrus.Info(time.Now().Format(time.RFC3339), " PostRecommendationsAnalyzeHandler start")

	if params.Content == nil { // possible?
		return specops.NewPostRecommendationsAnalyzeBadRequest().WithPayload(&specmodels.GenericError{Msg: "file is empty"})
	}

	img_barcode, err := common.ScanBarCodeFile(params.Content)
	if err != nil {
		return specops.NewPostRecommendationsAnalyzeBadRequest().WithPayload(&specmodels.GenericError{Msg: err.Error()})
	}

	logrus.Info(time.Now().Format(time.RFC3339), " PostRecommendationsAnalyzeHandler ", img_barcode, err)
	res, err := srv.levensteinRecommendations(img_barcode)
	if err != nil {
		return specops.NewPostRecommendationsAnalyzeBadRequest().WithPayload(&specmodels.GenericError{Msg: err.Error()})
	}

	payload := make([]*specmodels.Recommendation, len(res))
	for i := range payload {
		payload[i] = &specmodels.Recommendation{
			Articul:  res[i].Name,
			ShopName: res[i].ShopName,
			Barcode:  img_barcode,
			Price:    float64(res[i].Price),
			URL:      res[i].Url,
		}
	}
	return specops.NewPostRecommendationsAnalyzeOK().WithPayload(payload)
}

func (srv *RecommendatorService) GetPingHandler(params specops.GetPingParams) middleware.Responder {
	return specops.NewGetPingOK().WithPayload(&specmodels.Pong{Text: "service done!"})
}

func (srv *RecommendatorService) findByBarcode(barcode string) ([]common.Recommendation, error) {
	articul, err := common.GetProductArticul(barcode)
	if err != nil {
		return nil, err
	}

	// kazanexpressRecommendation := common.Recommendation{
	// 	Name:     "полотенце махровое",
	// 	ShopName: common.KazanExpressName,
	// 	Price:    5590,
	// 	Url:      "https://kazanexpress.ru",
	// }
	kazanexpressRecommendation, err := common.GetPriceFrom(common.KazanExpress, articul)
	if err != nil {
		logrus.Errorf("kazanExpress.getPricefrom err = %s", err)
		// return nil, err
	}
	ymRecommendation, err := common.GetPriceFrom(common.YandexMarket, articul)
	if err != nil {
		logrus.Errorf("yandexMarket.getPricefrom err = %s", err)
	}

	type product struct {
		ID       int    `db:"product_id"`
		Barcode  string `db:"barcode"`
		Articul  string `db:"articul"`
		ShopName string `db:"shop_name"`
		Price    int    `db:"barcode"`
	}
	selectQ := `select bp.barcode, p.id as product_id, p.articul, s.name
from barcode_products bp 
inner join products p on (bp.product_id = p.id)
inner join shops s on (p.shop_id = s.id)
where bp.barcode = ?`
	// selectQ := `select bd.barcode, p.url, from barcode_products bp inner join products p on (product_id = id) inner join shops s on (shop_id = id)`
	var products = make([]product, 0)
	if err = srv.db.Select(&products, selectQ, barcode); err != nil {
		logrus.Errorf("srv.db.Select products error = %s", err)
		return nil, err
	}
	// wrap tx
	for i := range products {
		p := products[i]
		// var newPrice int
		switch p.ShopName {
		case common.KazanExpressName:
			srv.updateProductInfo(p.ID, barcode, articul, kazanexpressRecommendation)
		case common.YandexMarketName:
			srv.updateProductInfo(p.ID, barcode, articul, ymRecommendation)
		default:
			logrus.Info("shop not found")
		}
	}

	res := make([]common.Recommendation, 2)
	res[0] = kazanexpressRecommendation
	res[1] = ymRecommendation
	return res, err
}

func (srv *RecommendatorService) levensteinRecommendations(barcode string) ([]common.Recommendation, error) {
	articul, err := common.GetProductArticul(barcode)
	if err != nil {
		return nil, err
	}
	//TODO: upstert

	// select all
	type product struct {
		ID int `db:"product_id"`
		// Barcode  string `db:"barcode"`
		Articul  string `db:"articul"`
		ShopName string `db:"shop_name"`
		Price    int    `db:"barcode"`
		URL      string `db:"url"`
	}
	selectQ := `select p.id as product_id, p.articul, s.name shop_name, p.url
from barcode_products bp 
inner join products p on (bp.product_id = p.id)
inner join shops s on (p.shop_id = s.id)
where bp.barcode = $1`
	// selectQ := `select bd.barcode, p.url, from barcode_products bp inner join products p on (product_id = id) inner join shops s on (shop_id = id)`
	var products = make([]product, 0)
	if err = srv.db.Select(&products, selectQ, barcode); err != nil {
		logrus.Errorf("srv.db.Select products barcode = %s error = %s", barcode, err)
		return nil, err
	}
	// calculate levenstein
	// productsDistance := make(map[int]int, len(products))    // расстояние до других товаров от текущего (id => расстояние)
	distancesProducts := make(map[int][]int, len(products)) // расстояние => товары с таким расстоянием
	productsMap := make(map[int]product, len(products))

	disctances := make([]int, 0, len(products))
	for _, p := range products {
		dist := leven.Distance(articul, p.Articul)
		distancesProducts[dist] = append(distancesProducts[dist], p.ID)
		// productsDistance[p.ID] = dist
		productsMap[p.ID] = p
		disctances = append(disctances, dist)
	}
	sort.Ints(disctances)
	// wrap tx
	recommendations := make([]common.Recommendation, 0, 5)
	distancesCnt := 5
	if len(disctances) < 5 {
		distancesCnt = len(disctances)
	}
	for _, dist := range disctances[:distancesCnt] {
		for _, pID := range distancesProducts[dist] {
			prod := productsMap[pID]
			recommendations = append(recommendations, common.Recommendation{
				Name:     prod.Articul,
				ShopName: prod.ShopName,
				Price:    prod.Price,
				Url:      prod.URL,
			})
		}
	}

	// var productDistances = make(map[productID]map[levensteinDist][]productID, 0)

	return recommendations, err
}

func (srv *RecommendatorService) updateProductInfo(productID int, barcode string, originalArticul string, rec common.Recommendation) error {
	newPrice := rec.Price
	// TODO: add updated_at
	updateQ := `update products p set p.price = ? where id = ?`
	res, err := srv.db.Exec(updateQ, newPrice, productID)
	if err != nil {
		logrus.Errorf("srv.db.Exec error = %s", err)
		return err // ?
	}
	if rowsCnt, _ := res.RowsAffected(); rowsCnt == 0 {
		var shopID string
		if err = srv.db.Get(&shopID, `select s.id from shops s where s.name = ?`, rec.ShopName); err != nil {
			logrus.Errorf("srv.db.Get shopName=%s error = %s", rec.ShopName, err)
			return err
		}
		insertQ := `insert into products (articul, url, shop_id, price) values (?, ?, ?, ?)`
		_, err := srv.db.Exec(insertQ, rec.Name, rec.Url, shopID, rec.Price)
		if err != nil {
			logrus.Errorf("srv.db.Exec articul=%s error = %s", rec.Name, err)
			return err // ?
		}
	}
	return nil
}
