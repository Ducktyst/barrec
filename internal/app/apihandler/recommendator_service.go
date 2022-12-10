package apihandler

import (
	"database/sql"
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/AlekSi/pointer"
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

	// 0 start
	res, err := srv.findByBarcode(params.Barcode)
	if err != nil {
		return specops.NewGetRecommendationsBarcodeBadRequest().WithPayload(&specmodels.GenericError{Msg: err.Error()})
	}
	// 0 end
	if err := srv.upsertProducts(res); err != nil {
		return specops.NewGetRecommendationsBarcodeBadRequest().WithPayload(&specmodels.GenericError{Msg: err.Error()})
	}

	// 1 start
	logrus.Info(time.Now().Format(time.RFC3339), " PostRecommendationsHandler levensteinRecommendations ", params.Barcode)
	res, err = srv.levensteinRecommendations(params.Barcode)
	if err != nil {
		return specops.NewPostRecommendationsAnalyzeBadRequest().WithPayload(&specmodels.GenericError{Msg: err.Error()})
	}
	logrus.Info(time.Now().Format(time.RFC3339), " PostRecommendationsHandler levensteinRecommendations end ", params.Barcode, err)
	// 1 end

	payload := make([]*specmodels.Recommendation, len(res))
	for i := range res {
		payload[i] = &specmodels.Recommendation{
			Articul:  pointer.ToString(res[i].Name),
			ShopName: pointer.ToString(res[i].ShopName),
			Barcode:  pointer.ToString(params.Barcode),
			Price:    pointer.ToInt64OrNil(res[i].Price),
			URL:      pointer.ToString(res[i].Url),
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

	// logrus.Info(time.Now().Format(time.RFC3339), " PostRecommendationsHandler findByBarcode start ", img_barcode)
	// _, err = srv.findByBarcode(img_barcode)
	// if err != nil {
	// 	return specops.NewPostRecommendationsAnalyzeBadRequest().WithPayload(&specmodels.GenericError{Msg: err.Error()})
	// }

	logrus.Info(time.Now().Format(time.RFC3339), " PostRecommendationsHandler levensteinRecommendations start ", img_barcode)
	res, err := srv.levensteinRecommendations(img_barcode)
	if err != nil {
		return specops.NewPostRecommendationsAnalyzeBadRequest().WithPayload(&specmodels.GenericError{Msg: err.Error()})
	}
	logrus.Info(time.Now().Format(time.RFC3339), " PostRecommendationsHandler levensteinRecommendations end ", img_barcode, err)

	payload := make([]*specmodels.Recommendation, len(res))
	for i := range res {
		payload[i] = &specmodels.Recommendation{
			Articul:  pointer.ToString(res[i].Name),
			ShopName: pointer.ToString(res[i].ShopName),
			Barcode:  pointer.ToString(img_barcode),
			Price:    pointer.ToInt64OrNil(res[i].Price),
			URL:      pointer.ToString(res[i].Url),
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
	for i := range res {
		payload[i] = &specmodels.Recommendation{
			Articul:  pointer.ToString(res[i].Name),
			ShopName: pointer.ToString(res[i].ShopName),
			Barcode:  pointer.ToString(img_barcode),
			Price:    pointer.ToInt64OrNil(res[i].Price),
			URL:      pointer.ToString(res[i].Url),
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
		logrus.Errorf("common.GetProductArticul error = %s", err)
		return nil, err
	}

	// kazanexpressRecommendation := common.Recommendation{
	// 	Name:     "полотенце махровое",
	// 	ShopName: common.KazanExpressName,
	// 	Price:    5590,
	// 	Url:      "https://kazanexpress.ru",
	// }
	kazanexpressRecommendation, err := common.GetPriceFrom(common.KazanExpress, barcode, articul)
	if err != nil {
		logrus.Errorf("kazanExpress.getPricefrom err = %s", err)
		// return nil, err
	}
	ymRecommendation, err := common.GetPriceFrom(common.YandexMarket, barcode, articul)
	if err != nil {
		logrus.Errorf("yandexMarket.getPricefrom err = %s", err)
	}

	// todo: зачем?
	// 	type product struct {
	// 		ID       int    `db:"product_id"`
	// 		Barcode  string `db:"barcode"`
	// 		Articul  string `db:"articul"`
	// 		ShopName string `db:"shop_name"`
	// 		Price    int    `db:"barcode"`
	// 	}
	// 	selectQ := `select bp.barcode, p.id as product_id, p.articul, s.name shop_name
	// from barcode_products bp
	// inner join products p on (bp.product_id = p.id)
	// inner join shops s on (p.shop_id = s.id)
	// where p.articul = $1`
	// 	// selectQ := `select bd.barcode, p.url, from barcode_products bp inner join products p on (product_id = id) inner join shops s on (shop_id = id)`
	// 	var products = make([]product, 0)
	// 	if err = srv.db.Select(&products, selectQ, articul); err != nil {
	// 		logrus.Errorf("srv.db.Select products articul = %s error = %s", articul, err)
	// 		return nil, err
	// 	}
	// 	// wrap tx

	res := make([]common.Recommendation, 2)
	res[0] = kazanexpressRecommendation
	res[1] = ymRecommendation
	return res, nil
}

func (srv *RecommendatorService) upsertProducts(res []common.Recommendation) error {
	for i := range res {
		rec := res[i]
		if err := srv.upsertProduct(rec.Barcode, rec.Name, rec); err != nil {
			return err
		}
	}
	return nil
}

func (srv *RecommendatorService) upsertProduct(barcode string, originalArticul string, rec common.Recommendation) error {
	logrus.Debugf("upsertProduct barcode = %s originalArticul = %s res.Price = %d rec.Name = %s rec.ShopName = %s", barcode, originalArticul, rec.Price, rec.Name, rec.ShopName)
	// TODO: add updated_at
	var shopCode string
	switch rec.ShopName {
	case common.KazanExpressName:
		shopCode = common.KazanExpressCode
	case common.YandexMarketName:
		shopCode = common.YandexMarketCode
	default:
		logrus.Errorf("shop not found: %s", rec.ShopName)
		return nil
	}

	var updateID int64
	selectQ := `select p.id from products p 
inner join barcode_products bp on (bp.product_id = p.id) 
inner join shops s on (s.id = p.shop_id)
where bp.barcode = $1 and s.name = $2 and p.articul = $3`
	err := srv.db.Get(&updateID, selectQ, rec.Barcode, shopCode, rec.Name)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logrus.Errorf("srv.upsertProduct srv.db.Get productID barcode = %s shopCode = %s rec.Name = %s error = %s", rec.Barcode, shopCode, rec.Name, err)
		return err
	}

	if updateID == 0 { // INSERT
		var shopID string
		if err = srv.db.Get(&shopID, `select s.id from shops s where s.name = $1`, shopCode); err != nil {
			logrus.Errorf("srv.upsertProduct getShopID shopName=%s error = %s", shopCode, err)
			return err
		}
		logrus.Infof("srv.upsertProduct shopID %s", shopID)

		// TODO: wrap transaction
		var newProductID int
		insertQ := `insert into products (articul, url, shop_id, price) values ($1, $2, $3, $4) returning id`
		err = srv.db.Get(&newProductID, insertQ, rec.Name, rec.Url, shopID, rec.Price)
		if err != nil {
			logrus.Errorf("srv.upsertProduct srv.db.Insert articul=%s error = %s", rec.Name, err)
			return err
		}

		insertQ = `insert into barcode_products (barcode, product_id) values ($1, $2)`
		_, err = srv.db.Exec(insertQ, barcode, newProductID)
		if err != nil {
			logrus.Errorf("srv.db.Insert barcode=%s product_id=%d error = %s", barcode, newProductID, err)
			return err
		}
		logrus.Infof("srv.upsertProduct insert barcode = %s articul = %s price = %d shopCode = %s, err = %w", barcode, rec.Name, rec.Price, shopCode)

	} else { // UPDATE
		updateQ := `update products set price = $1 where id = $2` // returning p.id не нужно, хватит и кол-ва измененных строк
		res, err := srv.db.Exec(updateQ, rec.Price, updateID)
		if err != nil {
			logrus.Errorf("srv.db.Exec error = %s", err)
			return err
		}
		logrus.Infof("srv.upsertProduct update barcode = %s articul = %s price = %d res = %v, err = %w", barcode, rec.Name, rec.Price, res, err)

		if _, err = res.RowsAffected(); err != nil {
			logrus.Errorf("srv.upsertProduct update rowsAffected error = %s", err)
			return err
		}
	}

	return nil
}

// levensteinRecommendations
// производит получение артикула товара по штрихкоду
// достает из бд все товары у которых есть штрихкод
// рассчитывает расстояние Левинштейна для каждого товара от имеющегося артикула
//
func (srv *RecommendatorService) levensteinRecommendations(barcode string) ([]common.Recommendation, error) {
	articul, err := common.GetProductArticul(barcode)
	if err != nil {
		return nil, fmt.Errorf("levensteinRecommendations GetProductArticul error: %w", err)
	}

	// select all
	type product struct {
		ID int `db:"product_id"`
		// Barcode  string `db:"barcode"`
		Articul  string `db:"articul"`
		ShopName string `db:"shop_name"`
		Barcode  string `db:"barcode"`
		Price    int    `db:"price"`
		URL      string `db:"url"`
	}
	selectQ := `select p.id as product_id, p.articul, s.name shop_name, bp.barcode, p.price price, p.url
from barcode_products bp 
inner join products p on (bp.product_id = p.id)
inner join shops s on (p.shop_id = s.id)`
	// selectQ := `select bd.barcode, p.url, from barcode_products bp inner join products p on (product_id = id) inner join shops s on (shop_id = id)`
	var products = make([]product, 0)
	if err = srv.db.Select(&products, selectQ); err != nil {
		logrus.Errorf("srv.db.Select products barcode = %s error = %s", barcode, err)
		return nil, err
	}
	// calculate levenstein
	distancesProducts := make(map[int][]int, len(products)) // расстояние => товары с таким расстоянием
	productsMap := make(map[int]product, len(products))

	distancesSet := make(map[int]struct{}, len(products))
	distances := make([]int, 0, len(products))
	for _, p := range products {
		p := p
		dist := leven.Distance(articul, p.Articul)
		distancesProducts[dist] = append(distancesProducts[dist], p.ID)
		productsMap[p.ID] = p

		// if dist < len([]rune(articul))-strings.Count(articul, " ") {
		// logrus.Info("dist < len(articul) ", dist, " < ", len([]rune(articul)))
		// logrus.Info("articul, p.Articul ", articul, ", ", p.Articul)
		if _, ok := distancesSet[dist]; !ok {
			distancesSet[dist] = struct{}{}
			distances = append(distances, dist)
		} // если нет совпадения даже по одной букве, то смысл добвлять слово
		// }
	}
	println()
	println("productsMap")
	println()
	println(productsMap)
	for id, p := range productsMap {
		println(id, p.Articul)
	}

	println("distancesProducts")
	for id, prods := range distancesProducts {
		println("dist = ", id)
		for id, pID := range prods {
			println(id, productsMap[pID].Articul)
		}
	}

	sort.Ints(distances)
	fmt.Println(distances)
	// wrap tx
	distancesCnt := 5
	recommendations := make([]common.Recommendation, 0, distancesCnt)
	if len(distances) < distancesCnt {
		distancesCnt = len(distances)
	}

	var found = 0
OUTER:
	for _, dist := range distances[:distancesCnt] {
		println("dist ", dist)
		for _, pID := range distancesProducts[dist] {
			prod := productsMap[pID]
			recommendations = append(recommendations, common.Recommendation{
				Name:     prod.Articul,
				ShopName: prod.ShopName,
				Price:    int64(prod.Price),
				Url:      prod.URL,
			})

			found++
			if found == distancesCnt {
				break OUTER
			}
		}
	}

	// logrus.Debug("articul ", articul)
	// logrus.Debug("productsMap", productsMap)
	// // logrus.Debug("distancesProducts ", distancesProducts)
	// logrus.Debug("disctances ", distances)
	// logrus.Debug("disctances Count ", distancesCnt)

	// var productDistances = make(map[productID]map[levensteinDist][]productID, 0)

	return recommendations, err
}
