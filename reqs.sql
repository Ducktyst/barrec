select p.id from products p 
inner join barcode_products bp on (bp.product_id = p.id) 
inner join shops s on (s.id = p.shop_id)
where bp.barcode =  '4690626047495' and s.name = 'kazanexpress' and p.articul = 'ЗАРЯДКА 1004C'