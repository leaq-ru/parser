package htmlfinder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	dress4carHTML = `<!DOCTYPE html>
<!--[if IE]><![endif]-->
<!--[if IE 8 ]><html prefix="og: http://ogp.me/ns# fb: http://ogp.me/ns/fb# business: http://ogp.me/ns/business#" dir="ltr" lang="ru" class="ie8"><![endif]-->
<!--[if IE 9 ]><html prefix="og: http://ogp.me/ns# fb: http://ogp.me/ns/fb# business: http://ogp.me/ns/business#" dir="ltr" lang="ru" class="ie9"><![endif]-->
<!--[if (gt IE 9)|!(IE)]><!-->
<html prefix="og: http://ogp.me/ns# fb: http://ogp.me/ns/fb# business: http://ogp.me/ns/business#" dir="ltr" lang="ru">
<!--<![endif]-->
<head>
    <meta charset="UTF-8"/>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>Авточехлы из экокожи, каркасные авточехлы в Нижнем Новгороде — цена, фото готовых работ</title>
    <base href="https://dress4car.ru/"/>
    <meta name="description" content="В нашей автостудии вы можете заказать авточехлы из экокожи по цене от 6000 руб, из жаккарда от 4500 руб, из экокожи в сочетании с жаккардом от 5800 руб, каркасные авточехлы с установкой по цене от 14000 руб. В Нижнем Новгороде, а также с доставкой по всей России"/>
    <link href="catalog/view/theme/next_default/stylesheet/stylesheet.css" rel="stylesheet">
    <link href="https://dress4car.ru/" rel="canonical"/>
    <link href="https://dress4car.ru/image/catalog/favicon.png" rel="icon"/>
</head>
<body class="common-home">
    <nav id="top">
        <div class="container">
            <div id="top-link" class="nav pull-left">
                <ul class="list-inline">
                    <li class="dropdown">
                        <a href="https://dress4car.ru/my-account/" title="Личный кабинет" class="dropdown-toggle" data-toggle="dropdown">
                            <i class="fa fa-user"></i>
                            <span class="hidden-xs hidden-sm">Личный кабинет</span>
                            <span class="caret"></span>
                        </a>
                        <ul class="dropdown-menu dropdown-menu-left">
                            <li>
                                <a href="https://dress4car.ru/create-account/">Регистрация</a>
                            </li>
                            <li>
                                <a href="https://dress4car.ru/login/">Авторизация</a>
                            </li>
                        </ul>
                        <div class="prmn-cmngr" data-confirm="false">
                            <div class="prmn-cmngr__content">
                                <div class="prmn-cmngr__title">
                                    <a class="prmn-cmngr__city">
                                        <span class="glyphicon glyphicon-map-marker"></span>
											г. Нижний Новгород
                                    </a>
                                </div>
                            </div>
                        </div>
                </ul>
            </div>
            <div id="top-links" class="nav pull-right">
                <ul class="list-inline">
                    <li>
                        <a href="https://dress4car.ru/wishlist/" id="wishlist-total" title="Мои закладки (0)">
                            <i class="fa fa-heart"></i>
                            <span class="hidden-xs hidden-sm hidden-md">Мои закладки (0)</span>
                        </a>
                    </li>
                    <li>
                        <a href="https://dress4car.ru/cart/" title="Корзина покупок">
                            <i class="fa fa-shopping-cart"></i>
                            <span class="hidden-xs hidden-sm hidden-md">Корзина покупок</span>
                        </a>
                    </li>
                    <li>
                        <a href="https://dress4car.ru/checkout/" title="Оформление заказа">
                            <i class="fa fa-share"></i>
                            <span class="hidden-xs hidden-sm hidden-md">Оформление заказа</span>
                        </a>
                    </li>
                </ul>
            </div>
        </div>
    </nav>
    <header>
        <div id="top2">
            <div class="container">
                <div class="row">
                    <div class="col-sm-3 vcenter w-100 text-center">
                        <div id="logo">
                            <img src="https://dress4car.ru/image/catalog/logo.png" title="Автостудия Dress4car" alt="Автостудия Dress4car"/>
                            <div class='data_name tagline'>
                                <span class="name tagline"></span>
                            </div>
                        </div>
                    </div>
                    <div class="col-sm-4 vcenter text-center">
                        <div class='phone'>
                            <a href="tel:+7 904 0555 202" onclick="yaCounter45326535.reachGoal('tel'); return true;">
                                <i class="fa fa-phone"></i>
                                <span class="">+7 904 0555 202</span>
                            </a>
                        </div>
                        <div class='data_open'>
                            <i class="fa fa-map-marker" aria-hidden="true"></i>
                            <span class="">
                                Нижний Новгород, ул. Дачная, д. 1-А 
                                <a href="/contact-us/" target="_blank">(Открыть карту)</a>
                            </span>
                        </div>
                    </div>
                    <div class="col-sm-4 vcenter">
                        <a href="/about_us" target="_blank" class="desc-btn">О компании</a>
                        <a href="/delivery" target="_blank" class="desc-btn">Оплата и доставка</a>
                        <a href="/info/" target="_blank" class="desc-btn">Материалы и цвета</a>
                        <a href="/otzyvy" target="_blank" class="desc-btn">Отзывы</a>
                        <a href="/contact-us/" target="_blank" class="desc-btn">Контакты</a>
                    </div>
                </div>
            </div>
        </div>
    </header>
    <div class="fix-menu">
        <div id="top3">
            <div class="container">
                <nav id="menu" class="navbar">
                    <div class="navbar-header">
                        <span id="category" class="visible-xs">Категории</span>
                        <button type="button" class="btn btn-navbar navbar-toggle" data-toggle="collapse" data-target=".navbar-ex1-collapse">
                            <i class="fa fa-bars"></i>
                        </button>
                    </div>
                    <div class="collapse navbar-collapse navbar-ex1-collapse">
                        <ul class="nav navbar-nav">
                            <li class="dropdown">
                                <a href="https://dress4car.ru/avtochekhly/" class="dropdown-toggle" data-toggle="dropdown">Авточехлы</a>
                                <div class="dropdown-menu">
                                    <div class="dropdown-inner">
                                        <ul class="list-unstyled">
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-chery/">Chery</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-chevrolet/">Chevrolet</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-daewoo/">Daewoo</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-datsun/">Datsun</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-ford/">Ford</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-gaz/">GAZ</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-geely/">Geely</a>
                                            </li>
                                        </ul>
                                        <ul class="list-unstyled">
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-great-wall/">Great Wall</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-honda/">Honda</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-hyundai/">Hyundai</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-kia/">KIA</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-lada/">Lada</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-mazda/">Mazda</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-mitsubishi/">Mitsubishi</a>
                                            </li>
                                        </ul>
                                        <ul class="list-unstyled">
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-nissan/">Nissan</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-opel/">Opel</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-peugeot/">Peugeot</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-ravon/">Ravon</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-renault/">Renault</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-skoda/">Skoda</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-ssangyong/">SsangYong</a>
                                            </li>
                                        </ul>
                                        <ul class="list-unstyled">
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-subaru/">Subaru</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-suzuki/">Suzuki</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-toyota/">Toyota</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-uaz/">UAZ</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-volkswagen/">Volkswagen</a>
                                            </li>
                                        </ul>
                                    </div>
                                    <a href="https://dress4car.ru/avtochekhly/" class="see-all">Показать все Авточехлы</a>
                                </div>
                            </li>
                            <li class="dropdown">
                                <a href="https://dress4car.ru/karkasnye-avtochekhly/" class="dropdown-toggle" data-toggle="dropdown">Каркасные авточехлы</a>
                                <div class="dropdown-menu">
                                    <div class="dropdown-inner">
                                        <ul class="list-unstyled">
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-chevrolet/">Chevrolet</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-ford/">Ford</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-honda/">Honda</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-hyundai/">Hyundai</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-kia/">KIA</a>
                                            </li>
                                        </ul>
                                        <ul class="list-unstyled">
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-lada/">Lada</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-land-rover/">Land Rover</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-mazda/">Mazda</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-mitsubishi/">Mitsubishi</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-nissan/">Nissan</a>
                                            </li>
                                        </ul>
                                        <ul class="list-unstyled">
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-opel/">Opel</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-renault/">Renault</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-skoda/">Skoda</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-ssangyong/">SsangYong</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-suzuki/">Suzuki</a>
                                            </li>
                                        </ul>
                                        <ul class="list-unstyled">
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-toyota/">Toyota</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-uaz/">UAZ</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-volkswagen/">Volkswagen</a>
                                            </li>
                                        </ul>
                                    </div>
                                    <a href="https://dress4car.ru/karkasnye-avtochekhly/" class="see-all">Показать все Каркасные авточехлы</a>
                                </div>
                            </li>
                            <li>
                                <a href="https://dress4car.ru/peretyazhka/">Перетяжка</a>
                            </li>
                            <li>
                                <a href="https://dress4car.ru/shumoizolyaciya/">Шумоизоляция</a>
                            </li>
                            <li>
                                <a href="https://dress4car.ru/avtoaksessuary/">Автоаксессуары</a>
                            </li>
                            <li>
                                <a href="/oformit-osago-onlajn" target="_blank">ОСАГО онлайн</a>
                            </li>
                        </ul>
                    </div>
                </nav>
            </div>
        </div>
    </div>
    <script src="catalog/view/javascript/jquery/jquery-2.1.1.min.js" type="text/javascript"></script>
    <script src="catalog/view/javascript/buyoneclick.js" type="text/javascript"></script>
    <script src="catalog/view/javascript/next-default/ls.js" type="text/javascript"></script>
    <script src="catalog/view/javascript/progroman/jquery.progroman.autocomplete.js" type="text/javascript"></script>
    <script src="catalog/view/javascript/progroman/jquery.progroman.city-manager.js" type="text/javascript"></script>
    <div class="container">
        <div class="row">
            <div id="content" class="col-sm-12">
                <div>
                    <div class="row">
                        <div class="main-page-card col-lg-5 col-md-5 col-sm-5 col-xs-11">
                            <div class="col-lg-5 col-md-6 col-sm-8 col-xs-11">
                                <a href="/avtochekhly/">
                                    <h3>Авточехлы из экокожи</h3>
                                </a>
                                <span>от 4500</span>
                                <ul>
                                    <li>
                                        <a href="/avtochekhly/2-chevrolet/">Chevrolet</a>
                                    </li>
                                    <li>
                                        <a href="/avtochekhly/2-ford/">Ford</a>
                                    </li>
                                    <li>
                                        <a href="/avtochekhly/2-hyundai/">Hyundai</a>
                                    </li>
                                    <li>
                                        <a href="/avtochekhly/2-kia/">KIA</a>
                                    </li>
                                    <li>
                                        <a href="/avtochekhly/2-renault/">Renault</a>
                                    </li>
                                    <li>
                                        <a href="/avtochekhly/2-volkswagen/">Volkswagen</a>
                                    </li>
                                    <li>
                                        <a href="/avtochekhly/">Другая марка</a>
                                    </li>
                                </ul>
                            </div>
                            <img data-src="/image/catalog/main-page/2.jpg" alt="Авточехлы из экокожи Dress4car" title="Авточехлы из экокожи Dress4car" class="h-100 lazy">
                            <a href="/avtochekhly/" class="desc-btn">Выбрать</a>
                        </div>
                        <div class="main-page-card col-lg-5 col-md-5 col-sm-5 col-xs-11">
                            <div class="col-lg-5 col-md-6 col-sm-8 col-xs-11">
                                <a href="/karkasnye-avtochekhly/">
                                    <h3>Каркасные авточехлы</h3>
                                </a>
                                <span>от 14000</span>
                                <ul>
                                    <li>
                                        <a href="/karkasnye-avtochekhly/1-chevrolet/">Chevrolet</a>
                                    </li>
                                    <li>
                                        <a href="/karkasnye-avtochekhly/1-ford/">Ford</a>
                                    </li>
                                    <li>
                                        <a href="/karkasnye-avtochekhly/1-hyundai/">Hyundai</a>
                                    </li>
                                    <li>
                                        <a href="/karkasnye-avtochekhly/1-kia/">KIA</a>
                                    </li>
                                    <li>
                                        <a href="/karkasnye-avtochekhly/1-mazda/">Mazda</a>
                                    </li>
                                    <li>
                                        <a href="/karkasnye-avtochekhly/1-volkswagen/">Volkswagen</a>
                                    </li>
                                    <li>
                                        <a href="/karkasnye-avtochekhly/">Другая марка</a>
                                    </li>
                                </ul>
                            </div>
                            <img data-src="/image/catalog/main-page/1.jpg" alt="Каркасные авточехлы Dress4car" title="Каркасные авточехлы Dress4car" class="h-100 lazy">
                            <a href="/karkasnye-avtochekhly/" class="desc-btn">Выбрать</a>
                        </div>
                        <div class="main-page-card col-lg-5 col-md-5 col-sm-5 col-xs-11">
                            <div class="col-lg-5 col-md-6 col-sm-8 col-xs-8">
                                <a href="/peretyazhka/">
                                    <h3>Перетяжка</h3>
                                </a>
                                <span>от 1000</span>
                                <ul>
                                    <li>
                                        <a href="/peretyazhka/dvernyh-kart">Дверные карты</a>
                                    </li>
                                    <li>
                                        <a href="/peretyazhka/rulya">Руль</a>
                                    </li>
                                    <li>
                                        <a href="/peretyazhka/rychaga-kpp">Рычаг и чехол КПП</a>
                                    </li>
                                    <li>
                                        <a href="/peretyazhka/centralnogo-podlokotnika">Центральный подлокотник</a>
                                    </li>
                                </ul>
                            </div>
                            <img data-src="/image/catalog/main-page/3.jpg" alt="Перетяжка салона от Dress4car" title="Перетяжка салона от Dress4car" class="w-100 lazy">
                            <a href="/peretyazhka/" class="desc-btn">Выбрать</a>
                        </div>
                        <div class="main-page-card col-lg-5 col-md-5 col-sm-5 col-xs-11">
                            <div class="col-lg-5 col-md-6 col-sm-8 col-xs-11">
                                <a href="/shumoizolyaciya/">
                                    <h3>Шумоизоляция</h3>
                                </a>
                                <span>от 6900</span>
                                <p>Шумоизоляция полная, или отдельно двери, пол, багажник</p>
                            </div>
                            <img data-src="/image/catalog/main-page/4.jpg" alt="Шумоизоляция автомобиля от Dress4car" title="Шумоизоляция автомобиля от Dress4car" class="w-100 lazy">
                            <a href="/shumoizolyaciya/" class="desc-btn">Выбрать</a>
                        </div>
                    </div>
                    <div class="main-page-desc">
                        <h2>Виды авточехлов</h2>
                        <h3>Модельные обычные</h3>
                        <p>Шьются индивидуально для каждой марки и модели авто, поэтому отлично сидят. Модельные это значит что сам чехол отличается в зависимости от марки-модели-комплектации авто, так как у всех авто сиденья разных форм и размеров. У модельных обычных чехлов нет спицевого каркаса. Они не пришиваются к заводской обивке, просто натягиваются. У нас в наличии лекала (значит можем сшить чехлы) на 99% современных авто.</p>
                        <p>
                            Готовые работы можно посмотреть и выбрать в разделе 
                            <a href="/avtochekhly/">Авточехлы</a>
                        </p>
                        <h3>Модельные каркасные</h3>
                        <p>
                            Тоже модельные, то есть точно по размерам именно для вашей марки-модели-комплектации авто, значит отлично сидят. Имеют вшитый спицевой каркас для лучшего повторения рельефа сидения. При установке закрепляются прошивом сквозь сиденье (нити идут сквозь обивку и «тело» сиденья и закрепляются с обратной стороны). Подробно процесс установки описан в 
                            <a href="/info/karkasnye-avtochehli/podrobnee" target="_blank">нашей статье</a>
                            .
                        </p>
                        <p>
                            Посмотреть готовые работы и выбрать можно по ссылке 
                            <a href="/karkasnye-avtochekhly/">Каркасные авточехлы</a>
                        </p>
                        <h3>Модельные обычные «с подшивом»</h3>
                        <p>Это чехлы «модельные обычные», только их «подшивают» к обивке сидения. Мы так не делаем, и не рекомендуем вам заказывать «с подшивом» в другом месте.</p>
                        <p>Суть «подшива»: иголкой и ниткой чуть зацепляется чехол за «родную» тканевую обивку, на глубине 2-3 мм, точечно в нескольких местах чехла. Соответственно такое крепление крайне ненадежно, и когда водитель и пассажиры сидят – кресло деформируется, и точки «подшива» тоже перемещаются относительно друг-друга/основы кресла/родной обивки.</p>
                        <p>В течение эксплуатации чехлов «с подшивом» родная обивка рвется и в местах подшива остаются дырки, затяжки и разрывы</p>
                        <h3>Универсальные</h3>
                        <p>Одинаковые для всех моделей авто, соответственно плохо сидят. Мы такие не продаем. Обычно можно встретить на авторынках</p>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div class="cta-news">
        <div class="container">
            <div class="row text-center">
                <div id="yandex_rtb_R-A-622298-1"></div>
                <script type="text/javascript">
                (function(w, d, n, s, t) {
                    w[n] = w[n] || [];
                    w[n].push(function() {
                        Ya.Context.AdvManager.render({
                            blockId: "R-A-622298-1",
                            renderTo: "yandex_rtb_R-A-622298-1",
                            async: true
                        });
                    });
                    t = d.getElementsByTagName("script")[0];
                    s = d.createElement("script");
                    s.type = "text/javascript";
                    s.src = "//an.yandex.ru/system/context.js";
                    s.async = true;
                    t.parentNode.insertBefore(s, t);
                })(this, this.document, "yandexContextAsyncCallbacks");
                </script>
            </div>
        </div>
    </div>
    <footer>
        <div id="alerts" class="container">
            <div class="row">
                <div class="col-sm-3">
                    <h4>Мы ВКонтакте</h4>
                    <div id="vk_groups"></div>
                </div>
                <div class="col-sm-3">
                    <h4>Информация</h4>
                    <ul class="list-unstyled">
                        <li>
                            <a href="/contact-us/">Контакты</a>
                        </li>
                        <li>
                            <a href="/info/">Материалы и цвета</a>
                        </li>
                        <li>
                            <a href="/otzyvy">Отзывы</a>
                        </li>
                        <li>
                            <a href="https://dress4car.ru/sitemap/">Карта сайта</a>
                        </li>
                        <li>
                            <a href="/franshiza-avtoatelye-dress4car/" target="_blank">Франшиза автоателье Dress4car</a>
                        </li>
                    </ul>
                </div>
                <div class="col-sm-3">
                    <h4>Личный кабинет</h4>
                    <ul class="list-unstyled">
                        <li>
                            <a href="https://dress4car.ru/my-account/">Личный кабинет</a>
                        </li>
                        <li>
                            <a href="https://dress4car.ru/order-history/">История заказов</a>
                        </li>
                        <li>
                            <a href="https://dress4car.ru/wishlist/">Мои закладки</a>
                        </li>
                        <li>
                            <a href="https://dress4car.ru/newsletter/">Рассылка новостей</a>
                        </li>
                    </ul>
                </div>
                <div class="col-sm-3">
                    <h4>Контакты</h4>
                    <div class="data-footer">
                        <p>
                            <i class="fa fa-phone"></i>
                            <a href="tel:+7 904 0555 202" onclick="yaCounter45326535.reachGoal('tel'); return true;">+7 904 0555 202</a>
                        </p>
                        <p>
                            <i class="fa fa-map-marker"></i>
                             Нижний Новгород, ул. Дачная, д. 1-А 
                            <a href="/contact-us/" target="_blank">(Открыть карту)</a>
                        </p>
                        <p class="short-description">Авточехлы из экокожи, каркасные авточехлы в Нижнем Новгороде</p>
                    </div>
                </div>
            </div>
            <hr>
            <p>Автостудия Dress4car © 2020</p>
            <div id='updown' class="hidden-xs">
                <button id="up" class='updown'>
                    <i class="fa fa-angle-up" aria-hidden="true"></i>
                </button>
                <button id="down" class='updown'>
                    <i class="fa fa-angle-down" aria-hidden="true"></i>
                </button>
            </div>
        </div>
        <span itemscope itemtype="http://schema.org/AutoPartsStore">
            <meta itemprop="name" content="Автостудия Dress4car"/>
            <link itemprop="url" href="https://dress4car.ru/"/>
            <link itemprop="image" href="https://dress4car.ru/image/catalog/logo.png"/>
            <meta itemprop="email" content="info@dress4car.ru"/>
            <meta itemprop="priceRange" content="RUB"/>
            <span itemprop="potentialAction" itemscope itemtype="http://schema.org/SearchAction">
                <meta itemprop="target" content="https://dress4car.ru/index.php?route=product/search&search={search_term_string}"/>
                <input type="hidden" itemprop="query-input" name="search_term_string">
            </span>
        </span>
    </footer>
    <script async src="catalog/view/javascript/next-default/lazy.js" type="text/javascript"></script>
    <link href="catalog/view/javascript/bootstrap/css/bootstrap.min.css" rel="stylesheet" media="screen"/>
    <link href="catalog/view/theme/next_default/stylesheet/full.css" rel="stylesheet">
    <link href="//fonts.googleapis.com/css?family=Open+Sans:400,700&amp;subset=cyrillic" rel="stylesheet" type="text/css"/>
    <script async src="catalog/view/javascript/bootstrap/js/bootstrap.min.js" type="text/javascript"></script>
    <script async src="catalog/view/javascript/next-default/common.js" type="text/javascript"></script>
    <link href="catalog/view/javascript/font-awesome/css/font-awesome.min.css" rel="stylesheet" type="text/css"/>
    <script async src="catalog/view/javascript/next-default/nextmenufix.js" type="text/javascript"></script>
    <script async src="catalog/view/javascript/next-default/jquery.easydropdown.js" type="text/javascript"></script>
    <script src="/otzyvy-vk/js/d_script.js"></script>
    <script type="text/javascript" src="https://vk.com/js/api/openapi.js?159"></script>
    <div id="vk_community_messages"></div>
    <script type="text/javascript">
    VK.Widgets.CommunityMessages("vk_community_messages", 144090016, {
        disableExpandChatSound: "1",
        disableNewMessagesSound: "1",
        tooltipButtonText: "Есть вопрос?"
    });
    VK.Widgets.Group("vk_groups", {
        mode: 3,
        no_cover: 1,
        color1: 'F7F7F7',
        color2: '2F2D38',
        color3: '3351A6'
    }, 144090016);
    </script>
    <script type="text/javascript">
    (function(d, w, c) {
        (w[c] = w[c] || []).push(function() {
            try {
                w.yaCounter45326535 = new Ya.Metrika({
                    id: 45326535,
                    clickmap: true,
                    trackLinks: true,
                    accurateTrackBounce: true,
                    webvisor: true
                });
            } catch (e) {}
        });
        var n = d.getElementsByTagName("script")[0],
            s = d.createElement("script"),
            f = function() {
                n.parentNode.insertBefore(s, n);
            };
        s.type = "text/javascript";
        s.async = true;
        s.src = "https://mc.yandex.ru/metrika/watch.js";
        if (w.opera == "[object Opera]") {
            d.addEventListener("DOMContentLoaded", f, false);
        } else {
            f();
        }
    })(document, window, "yandex_metrika_callbacks");
    </script>
    <noscript>
        <div>
            <img src="https://mc.yandex.ru/watch/45326535" style="position:absolute; left:-9999px;" alt=""/>
        </div>
    </noscript>
    <script async src="https://www.googletagmanager.com/gtag/js?id=UA-102581873-1"></script>
    <script>
    window.dataLayer = window.dataLayer || [];
    function gtag() {
        dataLayer.push(arguments);
    }
    gtag('js', new Date());

    gtag('config', 'UA-102581873-1');
    </script>
    <script type="text/javascript">
    !function() {
        var t = document.createElement("script");
        t.type = "text/javascript", t.async = !0, t.src = "https://vk.com/js/api/openapi.js?159", t.onload = function() {
            VK.Retargeting.Init("VK-RTRG-284712-5RVSH"), VK.Retargeting.Hit()
        }, document.head.appendChild(t)
    }();
    </script>
    <noscript>
        <img src="https://vk.com/rtrg?p=VK-RTRG-284712-5RVSH" style="position:fixed; left:-999px;" alt=""/>
    </noscript>
    <script async src="//yastatic.net/es5-shims/0.0.2/es5-shims.min.js"></script>
    <script async src="//yastatic.net/share2/share.js"></script>
    <script charset="UTF-8" src="https://partners.strahovkaru.ru/widget_constructor/constructor.js" type="text/javascript"></script>
    <script>
    (function() {
        if (!document.getElementById('strahovkaruru-widget'))
            return;
        var strahovkaruruWidgetObj = new strahovkaruruWidgetClass({
            "themeId": "2c805036-64c1-4af6-b39d-3b8228ad5d75"
        });
    })();
    </script>
    <div id="boc_order" class="modal fade">
        <div class="modal-dialog">
            <div class="modal-content">
                <form id="boc_form" action="" role="form">
                    <fieldset>
                        <div class="modal-header">
                            <button class="close" type="button" data-dismiss="modal">×</button>
                            <h2 id="boc_order_title" class="modal-title">Быстрый заказ</h2>
                        </div>
                        <div class="modal-body">
                            <div id="boc_product_field" class="col-xs-12"></div>
                            <div class="col-xs-12">
                                <div style="display:none">
                                    <input id="boc_admin_email" type="text" name="boc_admin_email" value="">
                                </div>
                                <div style="display:none">
                                    <input id="boc_product_id" type="text" name="boc_product_id">
                                </div>
                                <div class="input-group has-warning">
                                    <span class="input-group-addon">
                                        <i class="fa fa-fw fa-phone-square" aria-hidden="true"></i>
                                    </span>
                                    <input id="boc_phone" class="form-control required" type="tel" name="boc_phone" placeholder="Телефон" data-pattern="false">
                                </div>
                                <br/>
                            </div>
                            <div class="clearfix"></div>
                        </div>
                        <div class="modal-footer">
                            <div class="col-sm-2 hidden-xs"></div>
                            <div class="col-sm-8 col-xs-12">
                                <button type="submit" id="boc_submit" class="btn btn-lg btn-block btn-default" onclick="yaCounter45326535.reachGoal('fastorder'); return true;">Отправить</button>
                            </div>
                            <div class="col-sm-2 hidden-xs"></div>
                        </div>
                    </fieldset>
                </form>
            </div>
        </div>
    </div>
    <div id="boc_success" class="modal fade">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-body">
                    <div class="text-center">
                        <h4>
                            Спасибо за Ваш заказ!
                            <br/>
                            Мы свяжемся с Вами в самое ближайшее время.
                        </h4>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <script type="text/javascript">
    $('.boc_order_btn').on('click', function() {
        $.ajax({
            url: 'index.php?route=common/buyoneclick/info',
            type: 'post',
            data: $('#product input[type=\'text\'], #product input[type=\'hidden\'], #product input[type=\'radio\']:checked, #product input[type=\'checkbox\']:checked, #product select, #product textarea'),
            beforeSend: function() {
                $('.boc_order_btn').button('loading');
            },
            complete: function() {
                $('.boc_order_btn').button('reset');
            },
            success: function(data) {
                //console.log(data);
                $('#boc_product_field').html(data);
            },
            error: function(xhr, ajaxOptions, thrownError) {
                console.log(thrownError + "\r\n" + xhr.statusText + "\r\n" + xhr.responseText);
            }
        });
    });
    $('.boc_order_category_btn').on('click', function() {
        var for_post = {};
        for_post.product_id = $(this).attr('data-product_id');
        $.ajax({
            url: 'index.php?route=common/buyoneclick/info',
            type: 'post',
            data: for_post,
            beforeSend: function() {
                $('.boc_order_btn').button('loading');
            },
            complete: function() {
                $('.boc_order_btn').button('reset');
            },
            success: function(data) {
                //console.log(data);
                $('#boc_product_field').html(data);
            },
            error: function(xhr, ajaxOptions, thrownError) {
                console.log(thrownError + "\r\n" + xhr.statusText + "\r\n" + xhr.responseText);
            }
        });
    });
    </script>
</body>
</html>
`
	dress4carNoCityHTML = `<!DOCTYPE html>
<!--[if IE]><![endif]-->
<!--[if IE 8 ]><html prefix="og: http://ogp.me/ns# fb: http://ogp.me/ns/fb# business: http://ogp.me/ns/business#" dir="ltr" lang="ru" class="ie8"><![endif]-->
<!--[if IE 9 ]><html prefix="og: http://ogp.me/ns# fb: http://ogp.me/ns/fb# business: http://ogp.me/ns/business#" dir="ltr" lang="ru" class="ie9"><![endif]-->
<!--[if (gt IE 9)|!(IE)]><!-->
<html prefix="og: http://ogp.me/ns# fb: http://ogp.me/ns/fb# business: http://ogp.me/ns/business#" dir="ltr" lang="ru">
<!--<![endif]-->
<head>
    <meta charset="UTF-8"/>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>Авточехлы из экокожи, каркасные авточехлы в Нижнем Новгороде — цена, фото готовых работ</title>
    <base href="https://dress4car.ru/"/>
    <meta name="description" content="В нашей автостудии вы можете заказать авточехлы из экокожи по цене от 6000 руб, из жаккарда от 4500 руб, из экокожи в сочетании с жаккардом от 5800 руб, каркасные авточехлы с установкой по цене от 14000 руб. В Нижнем Новгороде, а также с доставкой по всей России"/>
    <link href="catalog/view/theme/next_default/stylesheet/stylesheet.css" rel="stylesheet">
    <link href="https://dress4car.ru/" rel="canonical"/>
    <link href="https://dress4car.ru/image/catalog/favicon.png" rel="icon"/>
</head>
<body class="common-home">
    <nav id="top">
        <div class="container">
            <div id="top-link" class="nav pull-left">
                <ul class="list-inline">
                    <li class="dropdown">
                        <a href="https://dress4car.ru/my-account/" title="Личный кабинет" class="dropdown-toggle" data-toggle="dropdown">
                            <i class="fa fa-user"></i>
                            <span class="hidden-xs hidden-sm">Личный кабинет</span>
                            <span class="caret"></span>
                        </a>
                        <ul class="dropdown-menu dropdown-menu-left">
                            <li>
                                <a href="https://dress4car.ru/create-account/">Регистрация</a>
                            </li>
                            <li>
                                <a href="https://dress4car.ru/login/">Авторизация</a>
                            </li>
                        </ul>
                        <div class="prmn-cmngr" data-confirm="false">
                            <div class="prmn-cmngr__content">
                                <div class="prmn-cmngr__title">
                                    <a class="prmn-cmngr__city">
                                        <span class="glyphicon glyphicon-map-marker"></span>
                                    </a>
                                </div>
                            </div>
                        </div>
                </ul>
            </div>
            <div id="top-links" class="nav pull-right">
                <ul class="list-inline">
                    <li>
                        <a href="https://dress4car.ru/wishlist/" id="wishlist-total" title="Мои закладки (0)">
                            <i class="fa fa-heart"></i>
                            <span class="hidden-xs hidden-sm hidden-md">Мои закладки (0)</span>
                        </a>
                    </li>
                    <li>
                        <a href="https://dress4car.ru/cart/" title="Корзина покупок">
                            <i class="fa fa-shopping-cart"></i>
                            <span class="hidden-xs hidden-sm hidden-md">Корзина покупок</span>
                        </a>
                    </li>
                    <li>
                        <a href="https://dress4car.ru/checkout/" title="Оформление заказа">
                            <i class="fa fa-share"></i>
                            <span class="hidden-xs hidden-sm hidden-md">Оформление заказа</span>
                        </a>
                    </li>
                </ul>
            </div>
        </div>
    </nav>
    <header>
        <div id="top2">
            <div class="container">
                <div class="row">
                    <div class="col-sm-3 vcenter w-100 text-center">
                        <div id="logo">
                            <img src="https://dress4car.ru/image/catalog/logo.png" title="Автостудия Dress4car" alt="Автостудия Dress4car"/>
                            <div class='data_name tagline'>
                                <span class="name tagline"></span>
                            </div>
                        </div>
                    </div>
                    <div class="col-sm-4 vcenter text-center">
                        <div class='phone'>
                            <a href="tel:+7 904 0555 202" onclick="yaCounter45326535.reachGoal('tel'); return true;">
                                <i class="fa fa-phone"></i>
                                <span class="">+7 904 0555 202</span>
                            </a>
                        </div>
                        <div class='data_open'>
                            <i class="fa fa-map-marker" aria-hidden="true"></i>
                            <span class="">
                                Нижний Новгород, ул. Дачная, д. 1-А 
                                <a href="/contact-us/" target="_blank">(Открыть карту)</a>
                            </span>
                        </div>
                    </div>
                    <div class="col-sm-4 vcenter">
                        <a href="/about_us" target="_blank" class="desc-btn">О компании</a>
                        <a href="/delivery" target="_blank" class="desc-btn">Оплата и доставка</a>
                        <a href="/info/" target="_blank" class="desc-btn">Материалы и цвета</a>
                        <a href="/otzyvy" target="_blank" class="desc-btn">Отзывы</a>
                        <a href="/contact-us/" target="_blank" class="desc-btn">Контакты</a>
                    </div>
                </div>
            </div>
        </div>
    </header>
    <div class="fix-menu">
        <div id="top3">
            <div class="container">
                <nav id="menu" class="navbar">
                    <div class="navbar-header">
                        <span id="category" class="visible-xs">Категории</span>
                        <button type="button" class="btn btn-navbar navbar-toggle" data-toggle="collapse" data-target=".navbar-ex1-collapse">
                            <i class="fa fa-bars"></i>
                        </button>
                    </div>
                    <div class="collapse navbar-collapse navbar-ex1-collapse">
                        <ul class="nav navbar-nav">
                            <li class="dropdown">
                                <a href="https://dress4car.ru/avtochekhly/" class="dropdown-toggle" data-toggle="dropdown">Авточехлы</a>
                                <div class="dropdown-menu">
                                    <div class="dropdown-inner">
                                        <ul class="list-unstyled">
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-chery/">Chery</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-chevrolet/">Chevrolet</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-daewoo/">Daewoo</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-datsun/">Datsun</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-ford/">Ford</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-gaz/">GAZ</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-geely/">Geely</a>
                                            </li>
                                        </ul>
                                        <ul class="list-unstyled">
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-great-wall/">Great Wall</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-honda/">Honda</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-hyundai/">Hyundai</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-kia/">KIA</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-lada/">Lada</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-mazda/">Mazda</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-mitsubishi/">Mitsubishi</a>
                                            </li>
                                        </ul>
                                        <ul class="list-unstyled">
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-nissan/">Nissan</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-opel/">Opel</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-peugeot/">Peugeot</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-ravon/">Ravon</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-renault/">Renault</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-skoda/">Skoda</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-ssangyong/">SsangYong</a>
                                            </li>
                                        </ul>
                                        <ul class="list-unstyled">
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-subaru/">Subaru</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-suzuki/">Suzuki</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-toyota/">Toyota</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-uaz/">UAZ</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/avtochekhly/2-volkswagen/">Volkswagen</a>
                                            </li>
                                        </ul>
                                    </div>
                                    <a href="https://dress4car.ru/avtochekhly/" class="see-all">Показать все Авточехлы</a>
                                </div>
                            </li>
                            <li class="dropdown">
                                <a href="https://dress4car.ru/karkasnye-avtochekhly/" class="dropdown-toggle" data-toggle="dropdown">Каркасные авточехлы</a>
                                <div class="dropdown-menu">
                                    <div class="dropdown-inner">
                                        <ul class="list-unstyled">
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-chevrolet/">Chevrolet</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-ford/">Ford</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-honda/">Honda</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-hyundai/">Hyundai</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-kia/">KIA</a>
                                            </li>
                                        </ul>
                                        <ul class="list-unstyled">
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-lada/">Lada</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-land-rover/">Land Rover</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-mazda/">Mazda</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-mitsubishi/">Mitsubishi</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-nissan/">Nissan</a>
                                            </li>
                                        </ul>
                                        <ul class="list-unstyled">
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-opel/">Opel</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-renault/">Renault</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-skoda/">Skoda</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-ssangyong/">SsangYong</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-suzuki/">Suzuki</a>
                                            </li>
                                        </ul>
                                        <ul class="list-unstyled">
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-toyota/">Toyota</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-uaz/">UAZ</a>
                                            </li>
                                            <li>
                                                <a href="https://dress4car.ru/karkasnye-avtochekhly/1-volkswagen/">Volkswagen</a>
                                            </li>
                                        </ul>
                                    </div>
                                    <a href="https://dress4car.ru/karkasnye-avtochekhly/" class="see-all">Показать все Каркасные авточехлы</a>
                                </div>
                            </li>
                            <li>
                                <a href="https://dress4car.ru/peretyazhka/">Перетяжка</a>
                            </li>
                            <li>
                                <a href="https://dress4car.ru/shumoizolyaciya/">Шумоизоляция</a>
                            </li>
                            <li>
                                <a href="https://dress4car.ru/avtoaksessuary/">Автоаксессуары</a>
                            </li>
                            <li>
                                <a href="/oformit-osago-onlajn" target="_blank">ОСАГО онлайн</a>
                            </li>
                        </ul>
                    </div>
                </nav>
            </div>
        </div>
    </div>
    <script src="catalog/view/javascript/jquery/jquery-2.1.1.min.js" type="text/javascript"></script>
    <script src="catalog/view/javascript/buyoneclick.js" type="text/javascript"></script>
    <script src="catalog/view/javascript/next-default/ls.js" type="text/javascript"></script>
    <script src="catalog/view/javascript/progroman/jquery.progroman.autocomplete.js" type="text/javascript"></script>
    <script src="catalog/view/javascript/progroman/jquery.progroman.city-manager.js" type="text/javascript"></script>
    <div class="container">
        <div class="row">
            <div id="content" class="col-sm-12">
                <div>
                    <div class="row">
                        <div class="main-page-card col-lg-5 col-md-5 col-sm-5 col-xs-11">
                            <div class="col-lg-5 col-md-6 col-sm-8 col-xs-11">
                                <a href="/avtochekhly/">
                                    <h3>Авточехлы из экокожи</h3>
                                </a>
                                <span>от 4500</span>
                                <ul>
                                    <li>
                                        <a href="/avtochekhly/2-chevrolet/">Chevrolet</a>
                                    </li>
                                    <li>
                                        <a href="/avtochekhly/2-ford/">Ford</a>
                                    </li>
                                    <li>
                                        <a href="/avtochekhly/2-hyundai/">Hyundai</a>
                                    </li>
                                    <li>
                                        <a href="/avtochekhly/2-kia/">KIA</a>
                                    </li>
                                    <li>
                                        <a href="/avtochekhly/2-renault/">Renault</a>
                                    </li>
                                    <li>
                                        <a href="/avtochekhly/2-volkswagen/">Volkswagen</a>
                                    </li>
                                    <li>
                                        <a href="/avtochekhly/">Другая марка</a>
                                    </li>
                                </ul>
                            </div>
                            <img data-src="/image/catalog/main-page/2.jpg" alt="Авточехлы из экокожи Dress4car" title="Авточехлы из экокожи Dress4car" class="h-100 lazy">
                            <a href="/avtochekhly/" class="desc-btn">Выбрать</a>
                        </div>
                        <div class="main-page-card col-lg-5 col-md-5 col-sm-5 col-xs-11">
                            <div class="col-lg-5 col-md-6 col-sm-8 col-xs-11">
                                <a href="/karkasnye-avtochekhly/">
                                    <h3>Каркасные авточехлы</h3>
                                </a>
                                <span>от 14000</span>
                                <ul>
                                    <li>
                                        <a href="/karkasnye-avtochekhly/1-chevrolet/">Chevrolet</a>
                                    </li>
                                    <li>
                                        <a href="/karkasnye-avtochekhly/1-ford/">Ford</a>
                                    </li>
                                    <li>
                                        <a href="/karkasnye-avtochekhly/1-hyundai/">Hyundai</a>
                                    </li>
                                    <li>
                                        <a href="/karkasnye-avtochekhly/1-kia/">KIA</a>
                                    </li>
                                    <li>
                                        <a href="/karkasnye-avtochekhly/1-mazda/">Mazda</a>
                                    </li>
                                    <li>
                                        <a href="/karkasnye-avtochekhly/1-volkswagen/">Volkswagen</a>
                                    </li>
                                    <li>
                                        <a href="/karkasnye-avtochekhly/">Другая марка</a>
                                    </li>
                                </ul>
                            </div>
                            <img data-src="/image/catalog/main-page/1.jpg" alt="Каркасные авточехлы Dress4car" title="Каркасные авточехлы Dress4car" class="h-100 lazy">
                            <a href="/karkasnye-avtochekhly/" class="desc-btn">Выбрать</a>
                        </div>
                        <div class="main-page-card col-lg-5 col-md-5 col-sm-5 col-xs-11">
                            <div class="col-lg-5 col-md-6 col-sm-8 col-xs-8">
                                <a href="/peretyazhka/">
                                    <h3>Перетяжка</h3>
                                </a>
                                <span>от 1000</span>
                                <ul>
                                    <li>
                                        <a href="/peretyazhka/dvernyh-kart">Дверные карты</a>
                                    </li>
                                    <li>
                                        <a href="/peretyazhka/rulya">Руль</a>
                                    </li>
                                    <li>
                                        <a href="/peretyazhka/rychaga-kpp">Рычаг и чехол КПП</a>
                                    </li>
                                    <li>
                                        <a href="/peretyazhka/centralnogo-podlokotnika">Центральный подлокотник</a>
                                    </li>
                                </ul>
                            </div>
                            <img data-src="/image/catalog/main-page/3.jpg" alt="Перетяжка салона от Dress4car" title="Перетяжка салона от Dress4car" class="w-100 lazy">
                            <a href="/peretyazhka/" class="desc-btn">Выбрать</a>
                        </div>
                        <div class="main-page-card col-lg-5 col-md-5 col-sm-5 col-xs-11">
                            <div class="col-lg-5 col-md-6 col-sm-8 col-xs-11">
                                <a href="/shumoizolyaciya/">
                                    <h3>Шумоизоляция</h3>
                                </a>
                                <span>от 6900</span>
                                <p>Шумоизоляция полная, или отдельно двери, пол, багажник</p>
                            </div>
                            <img data-src="/image/catalog/main-page/4.jpg" alt="Шумоизоляция автомобиля от Dress4car" title="Шумоизоляция автомобиля от Dress4car" class="w-100 lazy">
                            <a href="/shumoizolyaciya/" class="desc-btn">Выбрать</a>
                        </div>
                    </div>
                    <div class="main-page-desc">
                        <h2>Виды авточехлов</h2>
                        <h3>Модельные обычные</h3>
                        <p>Шьются индивидуально для каждой марки и модели авто, поэтому отлично сидят. Модельные это значит что сам чехол отличается в зависимости от марки-модели-комплектации авто, так как у всех авто сиденья разных форм и размеров. У модельных обычных чехлов нет спицевого каркаса. Они не пришиваются к заводской обивке, просто натягиваются. У нас в наличии лекала (значит можем сшить чехлы) на 99% современных авто.</p>
                        <p>
                            Готовые работы можно посмотреть и выбрать в разделе 
                            <a href="/avtochekhly/">Авточехлы</a>
                        </p>
                        <h3>Модельные каркасные</h3>
                        <p>
                            Тоже модельные, то есть точно по размерам именно для вашей марки-модели-комплектации авто, значит отлично сидят. Имеют вшитый спицевой каркас для лучшего повторения рельефа сидения. При установке закрепляются прошивом сквозь сиденье (нити идут сквозь обивку и «тело» сиденья и закрепляются с обратной стороны). Подробно процесс установки описан в 
                            <a href="/info/karkasnye-avtochehli/podrobnee" target="_blank">нашей статье</a>
                            .
                        </p>
                        <p>
                            Посмотреть готовые работы и выбрать можно по ссылке 
                            <a href="/karkasnye-avtochekhly/">Каркасные авточехлы</a>
                        </p>
                        <h3>Модельные обычные «с подшивом»</h3>
                        <p>Это чехлы «модельные обычные», только их «подшивают» к обивке сидения. Мы так не делаем, и не рекомендуем вам заказывать «с подшивом» в другом месте.</p>
                        <p>Суть «подшива»: иголкой и ниткой чуть зацепляется чехол за «родную» тканевую обивку, на глубине 2-3 мм, точечно в нескольких местах чехла. Соответственно такое крепление крайне ненадежно, и когда водитель и пассажиры сидят – кресло деформируется, и точки «подшива» тоже перемещаются относительно друг-друга/основы кресла/родной обивки.</p>
                        <p>В течение эксплуатации чехлов «с подшивом» родная обивка рвется и в местах подшива остаются дырки, затяжки и разрывы</p>
                        <h3>Универсальные</h3>
                        <p>Одинаковые для всех моделей авто, соответственно плохо сидят. Мы такие не продаем. Обычно можно встретить на авторынках</p>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div class="cta-news">
        <div class="container">
            <div class="row text-center">
                <div id="yandex_rtb_R-A-622298-1"></div>
                <script type="text/javascript">
                (function(w, d, n, s, t) {
                    w[n] = w[n] || [];
                    w[n].push(function() {
                        Ya.Context.AdvManager.render({
                            blockId: "R-A-622298-1",
                            renderTo: "yandex_rtb_R-A-622298-1",
                            async: true
                        });
                    });
                    t = d.getElementsByTagName("script")[0];
                    s = d.createElement("script");
                    s.type = "text/javascript";
                    s.src = "//an.yandex.ru/system/context.js";
                    s.async = true;
                    t.parentNode.insertBefore(s, t);
                })(this, this.document, "yandexContextAsyncCallbacks");
                </script>
            </div>
        </div>
    </div>
    <footer>
        <div id="alerts" class="container">
            <div class="row">
                <div class="col-sm-3">
                    <h4>Мы ВКонтакте</h4>
                    <div id="vk_groups"></div>
                </div>
                <div class="col-sm-3">
                    <h4>Информация</h4>
                    <ul class="list-unstyled">
                        <li>
                            <a href="/contact-us/">Контакты</a>
                        </li>
                        <li>
                            <a href="/info/">Материалы и цвета</a>
                        </li>
                        <li>
                            <a href="/otzyvy">Отзывы</a>
                        </li>
                        <li>
                            <a href="https://dress4car.ru/sitemap/">Карта сайта</a>
                        </li>
                        <li>
                            <a href="/franshiza-avtoatelye-dress4car/" target="_blank">Франшиза автоателье Dress4car</a>
                        </li>
                    </ul>
                </div>
                <div class="col-sm-3">
                    <h4>Личный кабинет</h4>
                    <ul class="list-unstyled">
                        <li>
                            <a href="https://dress4car.ru/my-account/">Личный кабинет</a>
                        </li>
                        <li>
                            <a href="https://dress4car.ru/order-history/">История заказов</a>
                        </li>
                        <li>
                            <a href="https://dress4car.ru/wishlist/">Мои закладки</a>
                        </li>
                        <li>
                            <a href="https://dress4car.ru/newsletter/">Рассылка новостей</a>
                        </li>
                    </ul>
                </div>
                <div class="col-sm-3">
                    <h4>Контакты</h4>
                    <div class="data-footer">
                        <p>
                            <i class="fa fa-phone"></i>
                            <a href="tel:+7 904 0555 202" onclick="yaCounter45326535.reachGoal('tel'); return true;">+7 904 0555 202</a>
                        </p>
                        <p>
                            <i class="fa fa-map-marker"></i>
                             Нижний Новгород, ул. Дачная, д. 1-А 
                            <a href="/contact-us/" target="_blank">(Открыть карту)</a>
                        </p>
                        <p class="short-description">Авточехлы из экокожи, каркасные авточехлы в Нижнем Новгороде</p>
                    </div>
                </div>
            </div>
            <hr>
            <p>Автостудия Dress4car © 2020</p>
            <div id='updown' class="hidden-xs">
                <button id="up" class='updown'>
                    <i class="fa fa-angle-up" aria-hidden="true"></i>
                </button>
                <button id="down" class='updown'>
                    <i class="fa fa-angle-down" aria-hidden="true"></i>
                </button>
            </div>
        </div>
        <span itemscope itemtype="http://schema.org/AutoPartsStore">
            <meta itemprop="name" content="Автостудия Dress4car"/>
            <link itemprop="url" href="https://dress4car.ru/"/>
            <link itemprop="image" href="https://dress4car.ru/image/catalog/logo.png"/>
            <meta itemprop="email" content="info@dress4car.ru"/>
            <meta itemprop="priceRange" content="RUB"/>
            <span itemprop="potentialAction" itemscope itemtype="http://schema.org/SearchAction">
                <meta itemprop="target" content="https://dress4car.ru/index.php?route=product/search&search={search_term_string}"/>
                <input type="hidden" itemprop="query-input" name="search_term_string">
            </span>
        </span>
    </footer>
    <script async src="catalog/view/javascript/next-default/lazy.js" type="text/javascript"></script>
    <link href="catalog/view/javascript/bootstrap/css/bootstrap.min.css" rel="stylesheet" media="screen"/>
    <link href="catalog/view/theme/next_default/stylesheet/full.css" rel="stylesheet">
    <link href="//fonts.googleapis.com/css?family=Open+Sans:400,700&amp;subset=cyrillic" rel="stylesheet" type="text/css"/>
    <script async src="catalog/view/javascript/bootstrap/js/bootstrap.min.js" type="text/javascript"></script>
    <script async src="catalog/view/javascript/next-default/common.js" type="text/javascript"></script>
    <link href="catalog/view/javascript/font-awesome/css/font-awesome.min.css" rel="stylesheet" type="text/css"/>
    <script async src="catalog/view/javascript/next-default/nextmenufix.js" type="text/javascript"></script>
    <script async src="catalog/view/javascript/next-default/jquery.easydropdown.js" type="text/javascript"></script>
    <script src="/otzyvy-vk/js/d_script.js"></script>
    <script type="text/javascript" src="https://vk.com/js/api/openapi.js?159"></script>
    <div id="vk_community_messages"></div>
    <script type="text/javascript">
    VK.Widgets.CommunityMessages("vk_community_messages", 144090016, {
        disableExpandChatSound: "1",
        disableNewMessagesSound: "1",
        tooltipButtonText: "Есть вопрос?"
    });
    VK.Widgets.Group("vk_groups", {
        mode: 3,
        no_cover: 1,
        color1: 'F7F7F7',
        color2: '2F2D38',
        color3: '3351A6'
    }, 144090016);
    </script>
    <script type="text/javascript">
    (function(d, w, c) {
        (w[c] = w[c] || []).push(function() {
            try {
                w.yaCounter45326535 = new Ya.Metrika({
                    id: 45326535,
                    clickmap: true,
                    trackLinks: true,
                    accurateTrackBounce: true,
                    webvisor: true
                });
            } catch (e) {}
        });
        var n = d.getElementsByTagName("script")[0],
            s = d.createElement("script"),
            f = function() {
                n.parentNode.insertBefore(s, n);
            };
        s.type = "text/javascript";
        s.async = true;
        s.src = "https://mc.yandex.ru/metrika/watch.js";
        if (w.opera == "[object Opera]") {
            d.addEventListener("DOMContentLoaded", f, false);
        } else {
            f();
        }
    })(document, window, "yandex_metrika_callbacks");
    </script>
    <noscript>
        <div>
            <img src="https://mc.yandex.ru/watch/45326535" style="position:absolute; left:-9999px;" alt=""/>
        </div>
    </noscript>
    <script async src="https://www.googletagmanager.com/gtag/js?id=UA-102581873-1"></script>
    <script>
    window.dataLayer = window.dataLayer || [];
    function gtag() {
        dataLayer.push(arguments);
    }
    gtag('js', new Date());

    gtag('config', 'UA-102581873-1');
    </script>
    <script type="text/javascript">
    !function() {
        var t = document.createElement("script");
        t.type = "text/javascript", t.async = !0, t.src = "https://vk.com/js/api/openapi.js?159", t.onload = function() {
            VK.Retargeting.Init("VK-RTRG-284712-5RVSH"), VK.Retargeting.Hit()
        }, document.head.appendChild(t)
    }();
    </script>
    <noscript>
        <img src="https://vk.com/rtrg?p=VK-RTRG-284712-5RVSH" style="position:fixed; left:-999px;" alt=""/>
    </noscript>
    <script async src="//yastatic.net/es5-shims/0.0.2/es5-shims.min.js"></script>
    <script async src="//yastatic.net/share2/share.js"></script>
    <script charset="UTF-8" src="https://partners.strahovkaru.ru/widget_constructor/constructor.js" type="text/javascript"></script>
    <script>
    (function() {
        if (!document.getElementById('strahovkaruru-widget'))
            return;
        var strahovkaruruWidgetObj = new strahovkaruruWidgetClass({
            "themeId": "2c805036-64c1-4af6-b39d-3b8228ad5d75"
        });
    })();
    </script>
    <div id="boc_order" class="modal fade">
        <div class="modal-dialog">
            <div class="modal-content">
                <form id="boc_form" action="" role="form">
                    <fieldset>
                        <div class="modal-header">
                            <button class="close" type="button" data-dismiss="modal">×</button>
                            <h2 id="boc_order_title" class="modal-title">Быстрый заказ</h2>
                        </div>
                        <div class="modal-body">
                            <div id="boc_product_field" class="col-xs-12"></div>
                            <div class="col-xs-12">
                                <div style="display:none">
                                    <input id="boc_admin_email" type="text" name="boc_admin_email" value="">
                                </div>
                                <div style="display:none">
                                    <input id="boc_product_id" type="text" name="boc_product_id">
                                </div>
                                <div class="input-group has-warning">
                                    <span class="input-group-addon">
                                        <i class="fa fa-fw fa-phone-square" aria-hidden="true"></i>
                                    </span>
                                    <input id="boc_phone" class="form-control required" type="tel" name="boc_phone" placeholder="Телефон" data-pattern="false">
                                </div>
                                <br/>
                            </div>
                            <div class="clearfix"></div>
                        </div>
                        <div class="modal-footer">
                            <div class="col-sm-2 hidden-xs"></div>
                            <div class="col-sm-8 col-xs-12">
                                <button type="submit" id="boc_submit" class="btn btn-lg btn-block btn-default" onclick="yaCounter45326535.reachGoal('fastorder'); return true;">Отправить</button>
                            </div>
                            <div class="col-sm-2 hidden-xs"></div>
                        </div>
                    </fieldset>
                </form>
            </div>
        </div>
    </div>
    <div id="boc_success" class="modal fade">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-body">
                    <div class="text-center">
                        <h4>
                            Спасибо за Ваш заказ!
                            <br/>
                            Мы свяжемся с Вами в самое ближайшее время.
                        </h4>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <script type="text/javascript">
    $('.boc_order_btn').on('click', function() {
        $.ajax({
            url: 'index.php?route=common/buyoneclick/info',
            type: 'post',
            data: $('#product input[type=\'text\'], #product input[type=\'hidden\'], #product input[type=\'radio\']:checked, #product input[type=\'checkbox\']:checked, #product select, #product textarea'),
            beforeSend: function() {
                $('.boc_order_btn').button('loading');
            },
            complete: function() {
                $('.boc_order_btn').button('reset');
            },
            success: function(data) {
                //console.log(data);
                $('#boc_product_field').html(data);
            },
            error: function(xhr, ajaxOptions, thrownError) {
                console.log(thrownError + "\r\n" + xhr.statusText + "\r\n" + xhr.responseText);
            }
        });
    });
    $('.boc_order_category_btn').on('click', function() {
        var for_post = {};
        for_post.product_id = $(this).attr('data-product_id');
        $.ajax({
            url: 'index.php?route=common/buyoneclick/info',
            type: 'post',
            data: for_post,
            beforeSend: function() {
                $('.boc_order_btn').button('loading');
            },
            complete: function() {
                $('.boc_order_btn').button('reset');
            },
            success: function(data) {
                //console.log(data);
                $('#boc_product_field').html(data);
            },
            error: function(xhr, ajaxOptions, thrownError) {
                console.log(thrownError + "\r\n" + xhr.statusText + "\r\n" + xhr.responseText);
            }
        });
    });
    </script>
</body>
</html>
`
)

func Benchmark_FindCity_found(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		FindCity(dress4carHTML)
	}
}

func Benchmark_FindCity_not_found(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		FindCity(dress4carNoCityHTML)
	}
}

func Test_FindCity_found(t *testing.T) {
	expectedResult := "Нижний Новгород"
	actualResult, actualIsFound := FindCity(dress4carHTML)

	assert.Equal(t, expectedResult, actualResult)
	assert.Equal(t, true, actualIsFound)
}

func Test_FindCity_not_found(t *testing.T) {
	actualResult, actualIsFound := FindCity(dress4carNoCityHTML)

	assert.Equal(t, "", actualResult)
	assert.Equal(t, false, actualIsFound)
}
