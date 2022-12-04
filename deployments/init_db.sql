insert into shops (name, go_search_template) values ('kazanexpress', '');
insert into shops (name, go_search_template) values ('yandexmarket', '');




insert into products(url, shop_id, price, articul) values ('https://kazanexpress.ru', 1, 5000, 'ИРИС СЛИВОЧНЫЙ"ЯШКИНО"140Г');
insert into products(url, shop_id, price, articul) values (
    'https://market.yandex.ru/product--iashkino-konfety-zagorskaia-slivochnaia-500gr/1754947307?nid=38605990&show-uid=16701619644219711840216001&context=search&text=%D0%98%D0%A0%D0%98%D0%A1%20%D0%A1%D0%9B%D0%98%D0%92%D0%9E%D0%A7%D0%9D%D0%AB%D0%99%22%D0%AF%D0%A8%D0%9A%D0%98%D0%9D%D0%9E%22140%D0%93%09&sku=101762657897&cpc=36t6AmG601Q-sz0DWPzNiMdhwCxefu1ak-GwZyvPkeMkc7cFnAg7ggCw10cqm9mvd7Ln5Wc0Fl_hURLuETwpZ_osf0zfv3u6a5ypm95o464FkCwy6KUnEwZi8Uetjxgpm-1MNRuBCj3kgCvrzTMxkyg0j9XH9_QgVhxPlC0iIKY%2C&do-waremd5=-A-iSx6wAqUT6vvmx_9aRw', 
    2, 19000, 'Яшкино Конфеты Загорская сливочная, 500гр');

insert into barcode_products(barcode, product_id) values ('4607015238693', 1);
insert into barcode_products(barcode, product_id) values ('4607015238693', 2);