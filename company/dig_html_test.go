package company

import (
	"context"
	"github.com/nnqq/scr-parser/call"
	"github.com/nnqq/scr-parser/mock"
	"github.com/nnqq/scr-proto/codegen/go/category"
	"github.com/nnqq/scr-proto/codegen/go/city"
	"github.com/stretchr/testify/assert"
	m "github.com/stretchr/testify/mock"
	"os"
	"testing"
)

func TestMain(tm *testing.M) {
	mockCity := &mock.City{}
	mockCity.
		On("Find", m.Anything, m.Anything, m.Anything).
		Return(&city.FindResponse{}, nil)
	call.City = mockCity

	mockCategory := &mock.Category{}
	mockCategory.
		On("Find", m.Anything, m.Anything, m.Anything).
		Return(&category.FindResponse{}, nil)
	call.Category = mockCategory

	os.Exit(tm.Run())
}

func BenchmarkDigHTML(b *testing.B) {
	dress4carHTML := []byte(`<!DOCTYPE html>
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
`)

	for i := 0; i < b.N; i += 1 {
		(&Company{}).digHTML(context.Background(), dress4carHTML, true, false, false)
	}
}

func TestDigHTML_dress4car(t *testing.T) {
	dress4carHTML := []byte(`<!DOCTYPE html>
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
`)

	expected := &Company{
		Title:       "Авточехлы из экокожи, каркасные авточехлы в Нижнем",
		Email:       "info@dress4car.ru",
		Phone:       79040555202,
		Description: "В нашей автостудии вы можете заказать авточехлы из экокожи по цене от 6000 руб, из жаккарда от 4500 руб, из экокожи в сочетании с жаккардом от 5800 руб, каркасные авточехлы с установкой по цене от 14000 руб. В Нижнем Новгороде, а также с доставкой по всей России",
	}
	actual := &Company{}
	actual.digHTML(context.Background(), dress4carHTML, true, false, false)

	assert.Equal(t, expected, actual)
}

func TestDigHTML_alarmSuzuki(t *testing.T) {
	alarmSuzukiHTML := []byte(`<!DOCTYPE html>
<html lang="ru">
<head>
  <meta charset="utf-8" />
<meta http-equiv="X-UA-Compatible" content="IE=edge" />
<meta content="Официальный дилерский центр Suzuki (Сузуки) в Санкт-Петербурге, автосалон «Аларм-Моторс SUZUKI», предлагает новые автомобили 2019 всего модельного ряда Suzuki, сервис и техническое обслуживание, а также оригинальные запчасти и аксессуары Suzuki." name="description" />
<meta content="официальный дилер сузуки спб" name="keywords" />
<meta name="viewport" content="width=device-width, initial-scale=1.0" />
<title>Официальный дилер Suzuki в Санкт-Петербурге</title>
<link rel="stylesheet" href="/css/vendor.min.css" />
<link rel="stylesheet" href="/css/main-mastersite.min.css" />

  
    <!-- Yandex.Metrika counter -->
<script type="text/javascript" >
(function(m,e,t,r,i,k,a){m[i]=m[i]||function(){(m[i].a=m[i].a||[]).push(arguments)};
m[i].l=1*new Date();k=e.createElement(t),a=e.getElementsByTagName(t)[0],k.async=1,k.src=r,a.parentNode.insertBefore(k,a)})
(window, document, "script", "https://mc.yandex.ru/metrika/tag.js", "ym");

ym(52744495, "init", {
clickmap:true,
trackLinks:true,
accurateTrackBounce:true,
webvisor:true
});
</script>
<noscript><div><img src="https://mc.yandex.ru/watch/52744495" style="position:absolute; left:-9999px;" alt="" /></div></noscript>
<!-- /Yandex.Metrika counter -->

<!-- Global site tag (gtag.js) - Google Analytics -->
<script async src="https://www.googletagmanager.com/gtag/js?id=UA-143018058-1"></script>
<script>
window.dataLayer = window.dataLayer || [];
function gtag(){dataLayer.push(arguments);}
gtag('js', new Date());

gtag('config', 'UA-143018058-1');
</script>

<script type="text/javascript" async src="//smartcallback.ru/api/SmartCallBack.js?t=AY5fUKecFuZ7viLHXqIF" charset="utf-8"></script>

<script async src="//callkeeper.ru/w/?38c3d9ae"></script>

<script src='https://code.reffection.com/pixel/tags/ec0d76f3-09fd-482d-958b-35d478fc480a'></script>
  
</head>
<body>
  
    <script>
        var Suzuki = window.Suzuki || {};
        Suzuki.dealerId = 'alarm-motors';
    </script>
  
    
    
  
  
	<div class="header-main-wrapper">
    	
	<header class="header">
		<div class="container">
			<div class="header__wrap">
				<div class="logo js-scroll-to-top">
					
						<div class="logo__symbol"><img src="/images/logo.svg" alt="Suzuki" /></div>
						<div class="logo__slogan"><img src="/images/slogan.svg" alt="Way of Life!" /></div>
					
				</div><div class="header__dealer">
					<div class="header__dealer-wrap container">
						<div class="header__dealer-info">
							<div class="header__dealer-name">Аларм-Моторс</div>
							<div class="header__dealer-desc">Официальный дилер Suzuki</div>
						</div>
						<button id="js-header-contact-button" class="header__contact-button"><svg class="header__contact-button-icon" width="24" height="24"><use xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="/images/icons/sprite.svg#icon-phone"></use></svg></button>
					</div>
					<div class="header__dealer-contacts">
						<div class="container">
							<div class="row justify-content-md-end"><div class="col-md header__dealer-item header__dealer-item_two">
									<div class="header__dealer-address">Санкт-Петербург, пр-т Маршала Жукова, 51</div><div class="header__dealer-phone"><a href="tel:&#43;7%20%28812%29%20606-77-59">&#43;7 (812) 606-77-59</a></div>
								</div><div class="col-md header__dealer-item header__dealer-item_two">
									<div class="header__dealer-address">Санкт-Петербург, Коломяжский пр-т, 18</div><div class="header__dealer-phone"><a href="tel:&#43;7%20%28812%29%20606-77-59">&#43;7 (812) 606-77-59</a></div>
								</div></div>
						</div>
					</div>						
				</div><button class="submenu-close"></button>
				<button id="js-header-burger" class="burger"></button>
			</div>
		</div>
		<div class="menu-wrap">
			<div class="container">
				<nav class="menu">
    <ul class="menu__list"><li class="menu__item"><a href="/auto/" class="menu__link">Модельный ряд</a><div class="submenu">
                <div class="submenu-model">
        <div class="container">
            <div class="submenu-model__info">
                <div class="submenu-model__title">Suzuki Vitara </div>
                <div class="submenu-model__price-block">
                    <div class="submenu-model__price">
                        <div class="submenu-model__price-label">цена от</div>
                        1 189 000 рублей
                    </div>
                    <div class="submenu-model__price">
                        <div class="submenu-model__price-label">кредит от</div>
                        12 862 рублей в месяц
                    </div>
                </div>
            </div>
        </div>
        <div class="submenu-model__pic"></div>
    </div>
    <div class="submenu-models-wrap">
        <div class="container">
            <div class="submenu-models"><a href="/auto/new-vitara/" class="submenu-models__thumbnail submenu-models__thumbnail_active " data-carname="Vitara " data-price="1 189 000  рублей" data-credit="12 862 рублей в месяц" data-image="/82b44b98-7608-4dfa-a8af-cd25edc0e90e/TM_Vitara_Motiv_23.jpg">
                    <div class="submenu-models__thumbnail-body">







<picture>
    <source srcset="/82b44b98-7608-4dfa-a8af-cd25edc0e90e/Vitara-1_hu805bac91f49ea07f5a4f0434fd6385df_74984_220x0_resize_q80_lanczos_2.png, /82b44b98-7608-4dfa-a8af-cd25edc0e90e/Vitara-1_hu805bac91f49ea07f5a4f0434fd6385df_74984_440x0_resize_q80_lanczos_2.png 2x">
    <img src="/82b44b98-7608-4dfa-a8af-cd25edc0e90e/Vitara-1_hu805bac91f49ea07f5a4f0434fd6385df_74984_220x0_resize_q80_lanczos_2.png" alt="" class="submenu-models__thumbnail-pic" />

</picture>
                        <div class="submenu-models__thumbnail-title">Vitara </div>
                        <div class="submenu-models__thumbnail-subtitle">цена от 1 189 000 рублей</div>
                        <div class="submenu-models__thumbnail-subtitle">кредит от 12 862 рублей в месяц</div>
                    </div>
                </a><a href="/auto/sx4-fl/" class="submenu-models__thumbnail " data-carname="SX4" data-price="1 289 000  рублей" data-credit="15 614 рублей в месяц" data-image="/82b44b98-7608-4dfa-a8af-cd25edc0e90e/S_Cross_hero_0035.jpg">
                    <div class="submenu-models__thumbnail-body">







<picture>
    <source srcset="/82b44b98-7608-4dfa-a8af-cd25edc0e90e/SX4_FL16_99999-BST03-041_10-2016%20GLX%20440x180_hu6808f9e354c3310d86ac438edffbfc85_104592_220x0_resize_q80_lanczos_2.png, /82b44b98-7608-4dfa-a8af-cd25edc0e90e/SX4_FL16_99999-BST03-041_10-2016%20GLX%20440x180_hu6808f9e354c3310d86ac438edffbfc85_104592_440x0_resize_q80_lanczos_2.png 2x">
    <img src="/82b44b98-7608-4dfa-a8af-cd25edc0e90e/SX4_FL16_99999-BST03-041_10-2016%20GLX%20440x180_hu6808f9e354c3310d86ac438edffbfc85_104592_220x0_resize_q80_lanczos_2.png" alt="" class="submenu-models__thumbnail-pic" />

</picture>
                        <div class="submenu-models__thumbnail-title">SX4</div>
                        <div class="submenu-models__thumbnail-subtitle">цена от 1 289 000 рублей</div>
                        <div class="submenu-models__thumbnail-subtitle">кредит от 15 614 рублей в месяц</div>
                    </div>
                </a><a href="/auto/jimny/" class="submenu-models__thumbnail " data-carname="all-new Jimny" data-price="1 639 000  рублей" data-credit="14 958 рублей в месяц" data-image="/82b44b98-7608-4dfa-a8af-cd25edc0e90e/Jimny_Ad_Photos_07-2018-018.jpg">
                    <div class="submenu-models__thumbnail-body">







<picture>
    <source srcset="/82b44b98-7608-4dfa-a8af-cd25edc0e90e/New%20Jimny-1_hu8643178ab76327f3eccafc9040efba70_59751_220x0_resize_q80_lanczos_2.png, /82b44b98-7608-4dfa-a8af-cd25edc0e90e/New%20Jimny-1_hu8643178ab76327f3eccafc9040efba70_59751_440x0_resize_q80_lanczos_2.png 2x">
    <img src="/82b44b98-7608-4dfa-a8af-cd25edc0e90e/New%20Jimny-1_hu8643178ab76327f3eccafc9040efba70_59751_220x0_resize_q80_lanczos_2.png" alt="" class="submenu-models__thumbnail-pic" />

</picture>
                        <div class="submenu-models__thumbnail-title">all-new Jimny</div>
                        <div class="submenu-models__thumbnail-subtitle">цена от 1 639 000 рублей</div>
                        <div class="submenu-models__thumbnail-subtitle">кредит от 14 958 рублей в месяц</div>
                    </div>
                </a>
            </div>
        </div>
    </div></div></li><li class="menu__item"><a href="/buy/" class="menu__link">Покупка</a><div class="submenu">
                <div class="submenu-block">
                    <div class="container">
                        <div class="row"><div class="offset-sm-2 col-sm-8 offset-md-0 col-md-4">
                                <div class="submenu__title">Новые автомобили</div>
                                <ul class="submenu__list">
                                <li class="submenu__item"><a href="/buy/new-auto/test-drajv/" class="submenu__link">Тест-драйв</a></li><li class="submenu__item"><a href="/buy/new-auto/credit-calc/" class="submenu__link">Рассчитать кредит</a></li><li class="submenu__item"><a href="/buy/new-auto/special-offers/" class="submenu__link">Cпециальные предложения</a></li><li class="submenu__item"><a href="/support/service-repair/loyalty-/" class="submenu__link">Программа лояльности</a></li><li class="submenu__item"><a href="/buy/new-auto/programma-semejnyj-avtomobil/" class="submenu__link">Семейный автомобиль</a></li><li class="submenu__item"><a href="/buy/new-auto/suzuki-insurance/" class="submenu__link">Страхование</a></li><li class="submenu__item"><a href="/buy/new-auto/ekspertnye-stati-o-suzuki-ot-oficialnogo-dilera/" class="submenu__link">Экспертные статьи</a></li></ul>
                            </div><div class="offset-sm-2 col-sm-8 offset-md-0 col-md-4">
                                <div class="submenu__title">Корпоративные продажи</div>
                                <ul class="submenu__list">
                                <li class="submenu__item"><a href="/buy/corporate/corporate-clients/" class="submenu__link">Программа корпоративных продаж</a></li></ul>
                            </div></div></div>
                </div></div></li><li class="menu__item"><a href="/support/" class="menu__link">Владельцам</a><div class="submenu">
                <div class="submenu-block">
                    <div class="container">
                        <div class="row"><div class="offset-sm-2 col-sm-8 offset-md-0 col-md-4">
                                <div class="submenu__title">Обслуживание и ремонт</div>
                                <ul class="submenu__list">
                                <li class="submenu__item"><a href="/suzuki/loyalty-/" class="submenu__link">Программа лояльности</a></li><li class="submenu__item"><a href="/support/service-repair/service/" class="submenu__link">Официальное обслуживание</a></li><li class="submenu__item"><a href="/support/service-repair/warranty-/" class="submenu__link">Гарантия</a></li><li class="submenu__item"><a href="/support/service-repair/assistance/" class="submenu__link">Suzuki Assistance</a></li><li class="submenu__item"><a href="/support/service-repair/advantages/" class="submenu__link">Suzuki привилегия 3&#43;</a></li></ul>
                            </div><div class="offset-sm-2 col-sm-8 offset-md-0 col-md-4">
                                <div class="submenu__title">Запчасти и аксессуары</div>
                                <ul class="submenu__list">
                                <li class="submenu__item"><a href="/support/spare-parts-accessories/spare-parts/" class="submenu__link">Оригинальные запчасти</a></li><li class="submenu__item"><a href="/support/spare-parts-accessories/oil/" class="submenu__link">Моторные масла</a></li><li class="submenu__item"><a href="/support/spare-parts-accessories/accessories/" class="submenu__link">Аксессуары</a></li><li class="submenu__item"><a href="/support/spare-parts-accessories/current/" class="submenu__link">Кузовные запчасти и ремонт</a></li><li class="submenu__item"><a href="/support/spare-parts-accessories/get-price/" class="submenu__link">Узнать стоимость детали</a></li></ul>
                            </div><div class="offset-sm-2 col-sm-8 offset-md-0 col-md-4">
                                <div class="submenu__title">Сервисные акции</div>
                                <ul class="submenu__list">
                                <li class="submenu__item"><a href="/support/service-actions/maslyanyj-servis/" class="submenu__link">Масляный сервис Suzuki</a></li><li class="submenu__item"><a href="/support/service-actions/servisnyj-komplekt-sceplenie/" class="submenu__link">Сервисные комплекты</a></li><li class="submenu__item"><a href="/support/service-actions/servisnye-nabory/" class="submenu__link">Сервисные наборы</a></li><li class="submenu__item"><a href="/support/service-actions/progressivnaya-skidka-v-zavisimosti-ot-vozrasta-a-m/" class="submenu__link">Прогрессивная скидка</a></li></ul>
                            </div></div><div class="buttons"><a class="buttons__item" href="/support/service-repair/zapis-na-to/">Запись на ТО</a><a class="buttons__item" href="/support/service-repair/prices/">Расcчитать ТО</a><a class="buttons__item" href="/support/service-repair/campaigns/">Сервисные кампании</a></div></div>
                </div></div></li><li class="menu__item"><a href="/buy/new-auto/special-offers/" class="menu__link">Акции</a></li><li class="menu__item"><a href="/about/o-kompanii/" class="menu__link">О компании</a></li><li class="menu__item"><a href="/about/contacts/" class="menu__link">Контакты</a></li><li class="menu__item"><a href="/press/news/" class="menu__link">Новости</a></li></ul>
</nav>
			</div>
		</div>
	</header>
	
    
    <main>
      
        
      
      
		
		<div class="hero">
			<div class="hero__slider">
            
                    
                
                    
                <div class="hero__slide" data-slide="0">
                    <div class="container">
                        <a href="/auto/new-vitara/" class="hero__link">







<picture>
    <source srcset="/images/f.png" data-srcset="/cc03c84d-76e4-44f6-86d0-76329d9a562b/HERO-desk-1800x805-F_VitA_hudb1f066827ca0666ab7497b2d995f1f7_748456_1800x0_resize_q75_lanczos.jpg" media="(min-width: 992px)">
    <source srcset="/images/f.png" data-srcset="/cc03c84d-76e4-44f6-86d0-76329d9a562b/HERO-plan-1600x1860_VitB_huc5978669f23b1f44a09e767e26f1e0b0_1550479_800x0_resize_q75_lanczos.jpg, /cc03c84d-76e4-44f6-86d0-76329d9a562b/HERO-plan-1600x1860_VitB_huc5978669f23b1f44a09e767e26f1e0b0_1550479_1600x0_resize_q75_lanczos.jpg 2x" media="(min-width: 768px)">
    <source srcset="/images/f.png" data-srcset="/cc03c84d-76e4-44f6-86d0-76329d9a562b/HERO-mob-960x1230_VitA_huc68f993a010f65dba895b61715bf42e7_629855_480x0_resize_q75_lanczos.jpg, /cc03c84d-76e4-44f6-86d0-76329d9a562b/HERO-mob-960x1230_VitA_huc68f993a010f65dba895b61715bf42e7_629855_960x0_resize_q75_lanczos.jpg 2x">
    <img src="/images/f.png" data-src="/cc03c84d-76e4-44f6-86d0-76329d9a562b/HERO-desk-1800x805-F_VitA_hudb1f066827ca0666ab7497b2d995f1f7_748456_1800x0_resize_q75_lanczos.jpg" alt="" class="hero__pic" />

</picture>
                        </a>
                        <div class="hero__slide-content">
                            <a href="/auto/new-vitara/" class="hero__link">
                                <h2>Vitara</h2>
                                <h3>Живи игрой!</h3>
                                
                            </a>
                            <a class="hero__button hero__button_border" href="/auto/new-vitara/">узнать больше</a>
                            <a class="hero__button" href="/buy/new-auto/special-offers/">СПЕЦПРЕДЛОЖЕНИЯ</a>
                        </div>
                        </div>
				</div>
            
                    
                
                    
                <div class="hero__slide" data-slide="1">
                    <div class="container">
                        <a href="/auto/sx4-fl/" class="hero__link">







<picture>
    <source srcset="/images/f.png" data-srcset="/a7ea5d63-20c6-4e5b-a642-77c967d0d578/HERO-desk-1800x805-F_SX4_hucc045de23c86f2a93b11d96f01035b81_758300_1800x0_resize_q75_lanczos.jpg" media="(min-width: 992px)">
    <source srcset="/images/f.png" data-srcset="/a7ea5d63-20c6-4e5b-a642-77c967d0d578/HERO-plan-1600x1860_SX4_hu35e00417c516022bb4ed0a94c2052d1b_1654430_800x0_resize_q75_lanczos.jpg, /a7ea5d63-20c6-4e5b-a642-77c967d0d578/HERO-plan-1600x1860_SX4_hu35e00417c516022bb4ed0a94c2052d1b_1654430_1600x0_resize_q75_lanczos.jpg 2x" media="(min-width: 768px)">
    <source srcset="/images/f.png" data-srcset="/a7ea5d63-20c6-4e5b-a642-77c967d0d578/HERO-mob-960x1230_SX4_hufb4e94d68aad9dbba98960f12a8bfb1d_532095_480x0_resize_q75_lanczos.jpg, /a7ea5d63-20c6-4e5b-a642-77c967d0d578/HERO-mob-960x1230_SX4_hufb4e94d68aad9dbba98960f12a8bfb1d_532095_960x0_resize_q75_lanczos.jpg 2x">
    <img src="/images/f.png" data-src="/a7ea5d63-20c6-4e5b-a642-77c967d0d578/HERO-desk-1800x805-F_SX4_hucc045de23c86f2a93b11d96f01035b81_758300_1800x0_resize_q75_lanczos.jpg" alt="" class="hero__pic" />

</picture>
                        </a>
                        <div class="hero__slide-content">
                            <a href="/auto/sx4-fl/" class="hero__link">
                                <h2>SX4</h2>
                                <h3>Не считая  километры</h3>
                                
                            </a>
                            <a class="hero__button hero__button_border" href="/auto/sx4-fl/">узнать больше</a>
                            <a class="hero__button" href="/buy/new-auto/special-offers/">СПЕЦПРЕДЛОЖЕНИЯ</a>
                        </div>
                        </div>
				</div>
            
                    
                
                    
                <div class="hero__slide" data-slide="2">
                    <div class="container">
                        <a href="/auto/jimny/" class="hero__link">







<picture>
    <source srcset="/images/f.png" data-srcset="/4e5571d6-06af-4105-892e-a3dcc96deed2/HERO-desk-1800x805-F_jimA_hu2790cdafe55b1454f77696939aadf5fe_599979_1800x0_resize_q75_lanczos.jpg" media="(min-width: 992px)">
    <source srcset="/images/f.png" data-srcset="/4e5571d6-06af-4105-892e-a3dcc96deed2/HERO-plan-1600x1860_JimB_hu3f1df895a79b235b0d32586b8871022d_1273787_800x0_resize_q75_lanczos.jpg, /4e5571d6-06af-4105-892e-a3dcc96deed2/HERO-plan-1600x1860_JimB_hu3f1df895a79b235b0d32586b8871022d_1273787_1600x0_resize_q75_lanczos.jpg 2x" media="(min-width: 768px)">
    <source srcset="/images/f.png" data-srcset="/4e5571d6-06af-4105-892e-a3dcc96deed2/HERO-mob-960x1230_JimA_hu0f4867904dbb5431c1cac9949073122c_454406_480x0_resize_q75_lanczos.jpg, /4e5571d6-06af-4105-892e-a3dcc96deed2/HERO-mob-960x1230_JimA_hu0f4867904dbb5431c1cac9949073122c_454406_960x0_resize_q75_lanczos.jpg 2x">
    <img src="/images/f.png" data-src="/4e5571d6-06af-4105-892e-a3dcc96deed2/HERO-desk-1800x805-F_jimA_hu2790cdafe55b1454f77696939aadf5fe_599979_1800x0_resize_q75_lanczos.jpg" alt="" class="hero__pic" />

</picture>
                        </a>
                        <div class="hero__slide-content">
                            <a href="/auto/jimny/" class="hero__link">
                                <h2>ALL-NEW JIMNY</h2>
                                <h3>Jimny. Такой один</h3>
                                
                            </a>
                            <a class="hero__button hero__button_border" href="/auto/jimny/">узнать больше</a>
                            <a class="hero__button" href="/buy/new-auto/special-offers/">СПЕЦПРЕДЛОЖЕНИЯ</a>
                        </div>
                        </div>
				</div>
            
                    
                <div class="hero__slide" data-slide="3">
                    <div class="container">
                        <a href="/suzuki/history/" class="hero__link">







<picture>
    <source srcset="/images/f.png" data-srcset="/e8e2a755-0757-45b6-963f-73fe81a7ed22/100-letie-suzuki-desk_hue60536acb5c8e0a784e24f9ec9d97705_1683666_1800x0_resize_q75_lanczos.jpg" media="(min-width: 992px)">
    <source srcset="/images/f.png" data-srcset="/e8e2a755-0757-45b6-963f-73fe81a7ed22/100-letie-suzuki-plan_hu342b2a2f5b76cae5a2bbf8e4bac8439a_1527283_800x0_resize_q75_lanczos.jpg, /e8e2a755-0757-45b6-963f-73fe81a7ed22/100-letie-suzuki-plan_hu342b2a2f5b76cae5a2bbf8e4bac8439a_1527283_1600x0_resize_q75_lanczos.jpg 2x" media="(min-width: 768px)">
    <source srcset="/images/f.png" data-srcset="/e8e2a755-0757-45b6-963f-73fe81a7ed22/100-letie-suzuki-mobil_hucb1d99706b7b18402e178fbe9374c718_731690_480x0_resize_q75_lanczos.jpg, /e8e2a755-0757-45b6-963f-73fe81a7ed22/100-letie-suzuki-mobil_hucb1d99706b7b18402e178fbe9374c718_731690_960x0_resize_q75_lanczos.jpg 2x">
    <img src="/images/f.png" data-src="/e8e2a755-0757-45b6-963f-73fe81a7ed22/100-letie-suzuki-desk_hue60536acb5c8e0a784e24f9ec9d97705_1683666_1800x0_resize_q75_lanczos.jpg" alt="" class="hero__pic" />

</picture>
                        </a>
                        <div class="hero__slide-content">
                            <a href="/suzuki/history/" class="hero__link">
                                
                                
                                
                            </a>
                            <a class="hero__button hero__button_border" href="/suzuki/history/">узнать больше</a>
                            
                        </div>
                        </div>
				</div>
            <div class="hero__slide" data-slide="4">
                    <div class="container">
                        <a href="/press/news/nashi-servisnye-centry-vozobnovlya-t-rabotu/" class="hero__link">







<picture>
    <source srcset="/images/f.png" data-srcset="/72772a11-ab95-436a-b2fd-bc149a4a6cf9/SERVICE_HERO-desk-1800x805_hueaae9de1be073179f75ffdfa02d662d1_868836_1800x0_resize_q75_lanczos.jpg" media="(min-width: 992px)">
    <source srcset="/images/f.png" data-srcset="/72772a11-ab95-436a-b2fd-bc149a4a6cf9/SERVICE_HERO-plan-1600x1860_hu7fe02caef4596fb66792987565c6a940_1925005_800x0_resize_q75_lanczos.jpg, /72772a11-ab95-436a-b2fd-bc149a4a6cf9/SERVICE_HERO-plan-1600x1860_hu7fe02caef4596fb66792987565c6a940_1925005_1600x0_resize_q75_lanczos.jpg 2x" media="(min-width: 768px)">
    <source srcset="/images/f.png" data-srcset="/72772a11-ab95-436a-b2fd-bc149a4a6cf9/SERVICE_HERO-mob-960x1230_hucf33fec5830e5d8d13450549ec873cb7_790928_480x0_resize_q75_lanczos.jpg, /72772a11-ab95-436a-b2fd-bc149a4a6cf9/SERVICE_HERO-mob-960x1230_hucf33fec5830e5d8d13450549ec873cb7_790928_960x0_resize_q75_lanczos.jpg 2x">
    <img src="/images/f.png" data-src="/72772a11-ab95-436a-b2fd-bc149a4a6cf9/SERVICE_HERO-desk-1800x805_hueaae9de1be073179f75ffdfa02d662d1_868836_1800x0_resize_q75_lanczos.jpg" alt="" class="hero__pic" />

</picture>
                        </a>
                        <div class="hero__slide-content">
                            <a href="/press/news/nashi-servisnye-centry-vozobnovlya-t-rabotu/" class="hero__link">
                                <h2>Ждем Вас</h2>
                                <h3>на сервис</h3>
                                
                            </a>
                            <a class="hero__button hero__button_border" href="/press/news/nashi-servisnye-centry-vozobnovlya-t-rabotu/">узнать больше</a>
                            
                        </div>
                        </div>
				</div>
            
            </div>
			<div class="hero__mouse-container">
				<div class="hero__mouse">
					<a href="#cars-slider"><img src="images/header-banner/mouse.svg" alt="mouse with arrows"></a>
				</div>
			</div>
			<div class="hero__slider-progress">
				<div class="hero__progress"></div>
			</div>
		</div>
		
        <div class="l-cars" id="cars-slider">
    <div class="slider-container">
        <div class="m-cars-slider center">
        
            <div class="m-cars-slider__slide cars-slider-active-slide" data-slide="0">
                <a href="/auto/new-vitara/">
                    <picture>
                        <source srcset="/82b44b98-7608-4dfa-a8af-cd25edc0e90e/780x720%20221_hu62b2a346cccb4f0b9084470fabdf9e0e_248805_700x0_resize_q80_lanczos_2.png">
                        <img class="m-car-photo" src="/82b44b98-7608-4dfa-a8af-cd25edc0e90e/780x720%20221_hu62b2a346cccb4f0b9084470fabdf9e0e_248805_351x0_resize_q80_lanczos_2.png" alt="Vitara ">
                    </picture>
                </a>
                <h3 class="slide-title col d-flex justify-content-center text-uppercase m-0"><a href="/auto/new-vitara/">Vitara </a></h3>
                <div class="slide-description">
                    <div class="price-container d-flex justify-content-center align-items-center">
                        <p class="price">цена от </p>
                        <p class="price-rub">1 189 000 рублей</p>
                    </div>
                    <div class="price-container d-flex justify-content-center align-items-center">
                        <p class="price">кредит от </p>
                        <p class="price-rub">12 862 рублей в месяц</p>
                    </div>
                    <div class="tools">
                        <a href="/auto/new-vitara/specifications/" class="tools__item"><svg class="tools__icon" width="23" height="23"><use xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="images/icons/sprite.svg#icon-characteristics"></use></svg><div class="tools__text">посмотреть<br> характеристики</div></a>
                    
                        <a href="/buy/new-auto/credit-calc/?model=novaya-vitara" class="tools__item"><svg class="tools__icon" width="23" height="23"><use xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="images/icons/sprite.svg#icon-credit"></use></svg><div class="tools__text">рассчитать<br> кредит</div></a>
                    </div>
                </div>
            </div>
        
            <div class="m-cars-slider__slide" data-slide="1">
                <a href="/auto/sx4-fl/">
                    <picture>
                        <source srcset="/82b44b98-7608-4dfa-a8af-cd25edc0e90e/780x720%20222_hua92d788153c4c7268d8177eabc463633_338552_700x0_resize_q80_lanczos_2.png">
                        <img class="m-car-photo" src="/82b44b98-7608-4dfa-a8af-cd25edc0e90e/780x720%20222_hua92d788153c4c7268d8177eabc463633_338552_351x0_resize_q80_lanczos_2.png" alt="SX4">
                    </picture>
                </a>
                <h3 class="slide-title col d-flex justify-content-center text-uppercase m-0"><a href="/auto/sx4-fl/">SX4</a></h3>
                <div class="slide-description">
                    <div class="price-container d-flex justify-content-center align-items-center">
                        <p class="price">цена от </p>
                        <p class="price-rub">1 289 000 рублей</p>
                    </div>
                    <div class="price-container d-flex justify-content-center align-items-center">
                        <p class="price">кредит от </p>
                        <p class="price-rub">15 614 рублей в месяц</p>
                    </div>
                    <div class="tools">
                        <a href="/auto/sx4-fl/specifications/" class="tools__item"><svg class="tools__icon" width="23" height="23"><use xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="images/icons/sprite.svg#icon-characteristics"></use></svg><div class="tools__text">посмотреть<br> характеристики</div></a>
                    
                        <a href="/buy/new-auto/credit-calc/?model=sx4-fl" class="tools__item"><svg class="tools__icon" width="23" height="23"><use xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="images/icons/sprite.svg#icon-credit"></use></svg><div class="tools__text">рассчитать<br> кредит</div></a>
                    </div>
                </div>
            </div>
        
            <div class="m-cars-slider__slide" data-slide="2">
                <a href="/auto/jimny/">
                    <picture>
                        <source srcset="/82b44b98-7608-4dfa-a8af-cd25edc0e90e/780x720%20223_hu3916b3b500b1ac4d5032b0f207f343e5_247573_700x0_resize_q80_lanczos_2.png">
                        <img class="m-car-photo" src="/82b44b98-7608-4dfa-a8af-cd25edc0e90e/780x720%20223_hu3916b3b500b1ac4d5032b0f207f343e5_247573_351x0_resize_q80_lanczos_2.png" alt="all-new Jimny">
                    </picture>
                </a>
                <h3 class="slide-title col d-flex justify-content-center text-uppercase m-0"><a href="/auto/jimny/">all-new Jimny</a></h3>
                <div class="slide-description">
                    <div class="price-container d-flex justify-content-center align-items-center">
                        <p class="price">цена от </p>
                        <p class="price-rub">1 639 000 рублей</p>
                    </div>
                    <div class="price-container d-flex justify-content-center align-items-center">
                        <p class="price">кредит от </p>
                        <p class="price-rub">14 958 рублей в месяц</p>
                    </div>
                    <div class="tools">
                        <a href="/auto/jimny/specifications/" class="tools__item"><svg class="tools__icon" width="23" height="23"><use xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="images/icons/sprite.svg#icon-characteristics"></use></svg><div class="tools__text">посмотреть<br> характеристики</div></a>
                    
                        <a href="/buy/new-auto/credit-calc/?model=all-new-jimny" class="tools__item"><svg class="tools__icon" width="23" height="23"><use xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="images/icons/sprite.svg#icon-credit"></use></svg><div class="tools__text">рассчитать<br> кредит</div></a>
                    </div>
                </div>
            </div>
        
            <div class="m-cars-slider__slide" data-slide="3">
                <a href="/auto/new-vitara/">
                    <picture>
                        <source srcset="/82b44b98-7608-4dfa-a8af-cd25edc0e90e/780x720%20221_hu62b2a346cccb4f0b9084470fabdf9e0e_248805_700x0_resize_q80_lanczos_2.png">
                        <img class="m-car-photo" src="/82b44b98-7608-4dfa-a8af-cd25edc0e90e/780x720%20221_hu62b2a346cccb4f0b9084470fabdf9e0e_248805_351x0_resize_q80_lanczos_2.png" alt="Vitara ">
                    </picture>
                </a>
                <h3 class="slide-title col d-flex justify-content-center text-uppercase m-0"><a href="/auto/new-vitara/">Vitara </a></h3>
                <div class="slide-description">
                    <div class="price-container d-flex justify-content-center align-items-center">
                        <p class="price">цена от </p>
                        <p class="price-rub">1 189 000 рублей</p>
                    </div>
                    <div class="price-container d-flex justify-content-center align-items-center">
                        <p class="price">кредит от </p>
                        <p class="price-rub">12 862 рублей в месяц</p>
                    </div>
                    <div class="tools">
                        <a href="/auto/new-vitara/specifications/" class="tools__item"><svg class="tools__icon" width="23" height="23"><use xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="images/icons/sprite.svg#icon-characteristics"></use></svg><div class="tools__text">посмотреть<br> характеристики</div></a>
                    
                        <a href="/buy/new-auto/credit-calc/?model=novaya-vitara" class="tools__item"><svg class="tools__icon" width="23" height="23"><use xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="images/icons/sprite.svg#icon-credit"></use></svg><div class="tools__text">рассчитать<br> кредит</div></a>
                    </div>
                </div>
            </div>
        
            <div class="m-cars-slider__slide" data-slide="4">
                <a href="/auto/sx4-fl/">
                    <picture>
                        <source srcset="/82b44b98-7608-4dfa-a8af-cd25edc0e90e/780x720%20222_hua92d788153c4c7268d8177eabc463633_338552_700x0_resize_q80_lanczos_2.png">
                        <img class="m-car-photo" src="/82b44b98-7608-4dfa-a8af-cd25edc0e90e/780x720%20222_hua92d788153c4c7268d8177eabc463633_338552_351x0_resize_q80_lanczos_2.png" alt="SX4">
                    </picture>
                </a>
                <h3 class="slide-title col d-flex justify-content-center text-uppercase m-0"><a href="/auto/sx4-fl/">SX4</a></h3>
                <div class="slide-description">
                    <div class="price-container d-flex justify-content-center align-items-center">
                        <p class="price">цена от </p>
                        <p class="price-rub">1 289 000 рублей</p>
                    </div>
                    <div class="price-container d-flex justify-content-center align-items-center">
                        <p class="price">кредит от </p>
                        <p class="price-rub">15 614 рублей в месяц</p>
                    </div>
                    <div class="tools">
                        <a href="/auto/sx4-fl/specifications/" class="tools__item"><svg class="tools__icon" width="23" height="23"><use xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="images/icons/sprite.svg#icon-characteristics"></use></svg><div class="tools__text">посмотреть<br> характеристики</div></a>
                    
                        <a href="/buy/new-auto/credit-calc/?model=sx4-fl" class="tools__item"><svg class="tools__icon" width="23" height="23"><use xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="images/icons/sprite.svg#icon-credit"></use></svg><div class="tools__text">рассчитать<br> кредит</div></a>
                    </div>
                </div>
            </div>
        
            <div class="m-cars-slider__slide" data-slide="5">
                <a href="/auto/jimny/">
                    <picture>
                        <source srcset="/82b44b98-7608-4dfa-a8af-cd25edc0e90e/780x720%20223_hu3916b3b500b1ac4d5032b0f207f343e5_247573_700x0_resize_q80_lanczos_2.png">
                        <img class="m-car-photo" src="/82b44b98-7608-4dfa-a8af-cd25edc0e90e/780x720%20223_hu3916b3b500b1ac4d5032b0f207f343e5_247573_351x0_resize_q80_lanczos_2.png" alt="all-new Jimny">
                    </picture>
                </a>
                <h3 class="slide-title col d-flex justify-content-center text-uppercase m-0"><a href="/auto/jimny/">all-new Jimny</a></h3>
                <div class="slide-description">
                    <div class="price-container d-flex justify-content-center align-items-center">
                        <p class="price">цена от </p>
                        <p class="price-rub">1 639 000 рублей</p>
                    </div>
                    <div class="price-container d-flex justify-content-center align-items-center">
                        <p class="price">кредит от </p>
                        <p class="price-rub">14 958 рублей в месяц</p>
                    </div>
                    <div class="tools">
                        <a href="/auto/jimny/specifications/" class="tools__item"><svg class="tools__icon" width="23" height="23"><use xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="images/icons/sprite.svg#icon-characteristics"></use></svg><div class="tools__text">посмотреть<br> характеристики</div></a>
                    
                        <a href="/buy/new-auto/credit-calc/?model=all-new-jimny" class="tools__item"><svg class="tools__icon" width="23" height="23"><use xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="images/icons/sprite.svg#icon-credit"></use></svg><div class="tools__text">рассчитать<br> кредит</div></a>
                    </div>
                </div>
            </div>
        
        </div>
    </div>
</div>

		
		<div class="for-owners">
			<div class="container">
				<div class="for-owners__title">ВЛАДЕЛЬЦАМ</div>
			</div>
			<div class="for-owners__slider">
			<div class="for-owners__slide">
					<a href="/support/service-repair/service/" class="for-owners__link">







<picture>
    <source srcset="/images/f.png" data-srcset="/11089874-e3af-4a45-9d40-cbc7af686dcc/Banner-desk-1600x495_hu691c4ca71fbcfae2650e9264556594c0_2317824_1600x0_resize_q75_lanczos.jpg" media="(min-width: 900px)">
    <source srcset="/images/f.png" data-srcset="/11089874-e3af-4a45-9d40-cbc7af686dcc/Banner-plan-1800x990_hudb941dc8606b05c6aa75f34f9f14d161_2876764_900x0_resize_q75_lanczos.jpg, /11089874-e3af-4a45-9d40-cbc7af686dcc/Banner-plan-1800x990_hudb941dc8606b05c6aa75f34f9f14d161_2876764_1800x0_resize_q75_lanczos.jpg 2x" media="(min-width: 600px)">
    <source srcset="/images/f.png" data-srcset="/11089874-e3af-4a45-9d40-cbc7af686dcc/Banner-mob-1200x990_huf13664f84f65c0836b3511278757a2c9_2615126_600x0_resize_q75_lanczos.jpg, /11089874-e3af-4a45-9d40-cbc7af686dcc/Banner-mob-1200x990_huf13664f84f65c0836b3511278757a2c9_2615126_1200x0_resize_q75_lanczos.jpg 2x">
    <img src="/images/f.png" data-src="/11089874-e3af-4a45-9d40-cbc7af686dcc/Banner-desk-1600x495_hu691c4ca71fbcfae2650e9264556594c0_2317824_1600x0_resize_q75_lanczos.jpg" alt="" class="lazyload for-owners__pic" />

</picture>
						<div class="for-owners__content">
							<div class="container">
								<div class="for-owners__slide-title">МЫ БЕРЕМ НА СЕБЯ<br />
ЗАБОТУ О ВАШЕМ SUZUKI</div>
								
							</div>
						</div>
					</a>
				</div>
			
			</div>
		</div>
		
		<div class="main-blocks-bg">
				<div class="quick-links">
					<div class="container">
						<div class="d-sm-flex justify-content-center">
                        <a href="#" class="quick-links__item">
								<svg class="quick-links__icon"><use xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="images/icons/sprite.svg#icon-rubl"></use></svg>
								<div>Стоимость то</div>
							</a>
                            <a href="/support/service-repair/zapis-na-to/" class="quick-links__item">
								<svg class="quick-links__icon"><use xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="images/icons/sprite.svg#icon-pen"></use></svg>
								<div>Запись на ТО</div>
							</a>
                            <a href="#" class="quick-links__item">
								<svg class="quick-links__icon"><use xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="images/icons/sprite.svg#icon-rubl"></use></svg>
								<div>Сервисные кампании</div>
							</a>
                            
						</div>
					</div>
				</div>
				
            <div class="news">
				<div class="container">
					<div class="news__block-title">НОВОСТИ</div>
					<div class="row news__slider">
                        
						<a href="/press/news/prodazhi-suzuki-v-yaponii-i-indii-rastut/" class="col-md-4 news__slide">
                                <div class="news__text">
                                    <div class="news__date">28 августа</div>
                                    <div class="news__year">2020</div>
                                    <div class="news__title">ПРОДАЖИ SUZUKI В ЯПОНИИ И ИНДИИ РАСТУТ</div>
                                </div>
                                <img class="news__pic" src="/d23ff0b2-13a0-4ce5-9b09-01c9ab22b01b/new-prev-28-08-20_hue27aa5bcf6e189cf268270a6ee35d2b8_418515_370x265_fill_q80_lanczos_smart1.jpg" alt="">
                                
                            </a>
                        
						<a href="/press/news/kazhdyj-tretij-vladelec-suzuki-vnov-vybiraet-avtomobil-etogo-brenda/" class="col-md-4 news__slide">
                                <div class="news__text">
                                    <div class="news__date">19 августа</div>
                                    <div class="news__year">2020</div>
                                    <div class="news__title">КАЖДЫЙ ТРЕТИЙ ВЛАДЕЛЕЦ SUZUKI ВНОВЬ ВЫБИРАЕТ АВТОМОБИЛЬ ЭТОГО БРЕНДА</div>
                                </div>
                                <img class="news__pic" src="/d23ff0b2-13a0-4ce5-9b09-01c9ab22b01b/vitara-prev-19-08-20_hu99c0122337756c475e3e200e9382bb8d_448864_370x265_fill_q80_lanczos_smart1.jpg" alt="">
                                
                            </a>
                        
						<a href="/press/news/suzuki-demonstiruet-polozhitelnu-dinamiku-prodazh/" class="col-md-4 news__slide">
                                <div class="news__text">
                                    <div class="news__date">7 августа</div>
                                    <div class="news__year">2020</div>
                                    <div class="news__title">SUZUKI демонстрирует положительную динамику продаж</div>
                                </div>
                                <img class="news__pic" src="/d23ff0b2-13a0-4ce5-9b09-01c9ab22b01b/prev-news_hu1c507d48a64239b06db5947cd3f3f92d_18360_370x265_fill_q80_lanczos_smart1.jpg" alt="">
                                
                            </a>
                        
                    </div>
                    <a href="/press/news/" class="button news__button">все новости</a>
				</div>
			</div>
    		</div>
		

      
        
      
    </main>
    
    
  </div>
	
  
	<footer class="footer">
		<div class="container">
			<div class="row footer__top">
				    <div class="col-sm-6 col-md-9">
        <div class="row"><div class="col-md-4">
                <div class="footer__sec-title footer-accordion-item">Мир Suzuki</div><ul class="footer__menu">
                    <li><a href="/suzuki/o-suzuki/">О Suzuki</a></li><li><a href="/suzuki/history/">История Suzuki</a></li><li><a href="/suzuki/loyalty-/">Программа лояльности</a></li></ul></div><div class="col-md-4">
                <div class="footer__sec-title footer-accordion-item">Пресс-центр</div><ul class="footer__menu">
                    <li><a href="/press/news/">Новости</a></li><li><a href="/press/our-events/">Наши мероприятия</a></li></ul></div><div class="col-md-4">
                <div class="footer__sec-title footer-accordion-item">О компании</div><ul class="footer__menu">
                    <li><a href="/about/o-kompanii/">О компании</a></li><li><a href="/about/contacts/">Контакты</a></li><li><a href="/about/-ridicheskaya-informaciya/">Юридическая информация</a></li></ul></div></div>
    </div>

				<div class="col-sm-6 col-md-3 footer__contact">
					<div class="footer__sec-title">связаться с нами</div><div class="footer__phone-email"><div class="footer__title">Аларм-Моторс Suzuki Юг</div><div class="footer__phone">+7 (812) 606-77-59</div><a href="mailto:online@alarm-motors.ru" class="footer__email">online@alarm-motors.ru</a>
					</div><div class="footer__phone-email"><div class="footer__title">Аларм-Моторс Suzuki Север</div><div class="footer__phone">+7 (812) 606-77-59</div><a href="mailto:online@alarm-motors.ru" class="footer__email">online@alarm-motors.ru</a>
					</div></div>
			</div>
			<div class="row footer__copy-media">
				<div class="col-sm-6 order-sm-1 col-lg-3">
					<div class="footer__media">
    <div class="footer__media-title">мы в соц. сетях</div>
    <div class="footer__media-content">
            <a class="footer__media-link" href="https://vk.com/alarmmotors_spb" target="_blank"><svg class="footer__media-icon"><use xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="/images/icons/sprite.svg#icon-vk"></use></svg></a>
            <a class="footer__media-link" href="https://www.instagram.com/alarmmotors.ru/" target="_blank"><svg class="footer__media-icon"><use xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="/images/icons/sprite.svg#icon-instagram"></use></svg></a>
    </div>
</div>

				</div>
				<div class="col-sm-6 align-self-sm-center col-lg-9">
					<div class="footer__copy">
						&copy; 2020 <br class="d-none-md">Аларм-Моторс
					</div>
				</div>
			</div>
		</div>
	</footer>
	
  <script src="/js/vendor.min.js" type="text/javascript"></script>
<script src="/js/main-mastersite.min.js" type="text/javascript"></script>


  <script type="text/javascript" src="//stock.cardigital.ru/load?id=3585"></script>
</body>
</html>
`)

	expected := &Company{
		Title:       "Официальный дилер Suzuki в Санкт-Петербурге",
		Email:       "online@alarm-motors.ru",
		Phone:       78126067759,
		Description: "Официальный дилерский центр Suzuki (Сузуки) в Санкт-Петербурге, автосалон «Аларм-Моторс SUZUKI», предлагает новые автомобили 2019 всего модельного ряда Suzuki, сервис и техническое обслуживание, а также оригинальные запчасти и аксессуары Suzuki.",
		Social: &social{
			Instagram: &item{URL: "https://www.instagram.com/alarmmotors.ru/"},
		},
	}
	actual := &Company{}
	actual.digHTML(context.Background(), alarmSuzukiHTML, true, false, false)

	assert.Equal(t, expected, actual)
}

func TestDigHTML_news50ru(t *testing.T) {
	news050ruHTML := []byte(`<!doctype html>
<html lang="ru">
<head>
	<meta charset="windows-1251" />
	<title>Новости 0-50.ru | </title>
	<meta name="description" content="">
	<meta name="author" content="Denis V. Korolkov right.ked@gmail.com">
	<!--[if lt IE 9]>
		<script src="http://html5shim.googlecode.com/svn/trunk/html5.js"></script>
	<![endif]-->
	<meta name="rp51afcaa28a694b82948bd3a36a0490f8" content="f3aeec7e35993bbadcd70e9a9c5bb26d" />
	<!-- Mobile Specific Metas -->
	<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1" /> 
	
	<script type="text/javascript">
		var browser			= navigator.userAgent;
		var browserRegex	= /(Android|BlackBerry|IEMobile|Nokia|iP(ad|hone|od)|Opera M(obi|ini))/;
		var isMobile		= false;
		if(browser.match(browserRegex)) {
			isMobile			= true;
			addEventListener("load", function() { setTimeout(hideURLbar, 0); }, false);
			function hideURLbar(){
				window.scrollTo(0,1);
			}
		}
	</script>
	
	<!-- CSS -->
	<link rel="stylesheet" href="/includes/base.css">
	<link rel="stylesheet" href="/includes/style.css">
	<link rel="stylesheet" href="/includes/layout.css">
	<link rel="stylesheet" href="/includes/amazium.css">
	<!-- Favicons -->
	<link rel="shortcut icon" href="/images/favicon.ico">

	<!-- To Top scripts -->
	<script src="/includes/smoothscroll.js"type="text/javascript" ></script>
	<script src="http://code.jquery.com/jquery-1.9.1.min.js" type="text/javascript"></script>
	<script src="/includes/jquery.easing.1.3.js" type="text/javascript"></script>
	<script src="/includes/jquery.ui.totop.js" type="text/javascript"></script>
	<script type="text/javascript">
		$(document).ready(function() {
			$().UItoTop({ easingType: 'easeOutQuart',text: 'Наверх' });
		});
	</script>
	
<!-- Standard iPhone, iPod touch -->
<link rel="apple-touch-icon" sizes="57x57" href="/apple-touch-icon-57.png" />
<!-- Retina iPhone, iPod touch -->
<link rel="apple-touch-icon" sizes="114x114" href="/apple-touch-icon-114.png" />
<!-- Standard iPad -->
<link rel="apple-touch-icon" sizes="72x72" href="/apple-touch-icon-72.png" />
<!-- Retina iPad -->
<link rel="apple-touch-icon" sizes="144x144" href="/apple-touch-icon-144.png" />

<script type="text/javascript" src="/include/swfobject/swfobject.js"></script>	
</head>
<body>
<!-- Yandex.Metrika counter -->
<script type="text/javascript">
(function (d, w, c) {
    (w[c] = w[c] || []).push(function() {
        try {
            w.yaCounter21009928 = new Ya.Metrika({id:21009928,
                    webvisor:true,
                    clickmap:true,
                    trackLinks:true,
                    accurateTrackBounce:true});
        } catch(e) { }
    });

    var n = d.getElementsByTagName("script")[0],
        s = d.createElement("script"),
        f = function () { n.parentNode.insertBefore(s, n); };
    s.type = "text/javascript";
    s.async = true;
    s.src = (d.location.protocol == "https:" ? "https:" : "http:") + "//mc.yandex.ru/metrika/watch.js";

    if (w.opera == "[object Opera]") {
        d.addEventListener("DOMContentLoaded", f, false);
    } else { f(); }
})(document, window, "yandex_metrika_callbacks");
</script>
<noscript><div><img src="//mc.yandex.ru/watch/21009928" style="position:absolute; left:-9999px;" alt="" /></div></noscript>
<!-- /Yandex.Metrika counter -->
<script async="async" src="https://w.uptolike.com/widgets/v1/zp.js?pid=46593"></script>

<header>
	<div class="row hide-phone">
		<div class="grid_12" id="label_top"><ul class="nomargin nobullet"> | <li class="nobullet inline noindent"><a href="/news/label/%CF%EE%E3%EE%E4%E0+%E2+%C5%EA%E0%F2%E5%F0%E8%ED%E1%F3%F0%E3%E5+%E8+%D1%E2%E5%F0%E4%EB%EE%E2%F1%EA%EE%E9+%EE%E1%EB%E0%F1%F2%E8">Погода в Екатеринбурге и Свердловской области</a> | </li><li class="nobullet inline noindent"><a href="/news/label/%D7%F2%EE+%EF%F0%E8%E3%EE%F2%EE%E2%E8%F2%FC+%ED%E0+%F3%E6%E8%ED+%F0%E5%F6%E5%EF%F2+%F1+%F4%EE%F2%EE">Что приготовить на ужин рецепт с фото</a> | </li><li class="nobullet inline noindent"><a href="/news/label/%CD%EE%E2%EE%F1%F2%E8+%C5%EA%E0%F2%E5%F0%E8%ED%E1%F3%F0%E3%E0">Новости Екатеринбурга</a> | </li><li class="nobullet inline noindent"><a href="/news/label/96women.ru+-+%C6%E8%E2%EE%E9+%C6%E5%ED%F1%EA%E8%E9+%C6%F3%F0%ED%E0%EB">96women.ru - Живой Женский Журнал</a> | </li></ul></div>
	</div>
	<div class="row hide-phone">
		<nav class="grid_12" id="link" >
			<ul class="float-right">
				<li><a href="/"><span>Главная</span></a></li>
				<li><a href="/editorial.html"><span>Редакция</span></a></li><li><a href="/advertising.html"><span>Реклама</span></a></li><li><a href="/aboutus.html"><span>О проекте</span></a></li>
				<li><a href="/rss/"><span>RSS</span></a></li>
                <li><a href="/sendmail/"><span>Обратная связь</span></a></li>
				<li><a href="/news/"><span>Все новости</span></a></li>
			</ul>
		</nav>
	</div>
	<div class="row">
		<div class="grid_2" id="logo"><a href="http://0-50.ru/"><img src="/images/logo.png" alt="0-50.RU" title="0-50.RU" width="140" height="88" /></a></div>
		<div class="grid_10 hide-phone row" id="top_banner">
			<div class="grid_6"></div>
			<div class="grid_2"></div>
			<div class="grid_2"><div class="banner" id="banner_26"><a href="/banners/26/" title="" target="_blank" ><img src="/images/banners/73049_original.jpg" /></a></div></div>
		</div>
	</div>
	<div class="row hide-phone">
	<nav class="grid_12" id="news_menu">
			<ul class="nomargin nobullet">
				<li class="nobullet inline"><a href="/news/eburg/">Екатеринбург</a></li>
				<li class="nobullet inline"><a href="/news/russia/">Россия и мир</a></li>
				<li class="nobullet inline"><a href="/news/education/">Образование</a></li>
				<li class="nobullet inline"><a href="/news/polit/">Недвижимость</a></li>
				<li class="nobullet inline"><a href="/news/health/">Здоровье</a></li>
				<li class="nobullet inline"><a href="/news/sport/">Спорт</a></li>
				<li class="nobullet inline"><a href="/news/incident/">Происшествия</a></li>
				<li class="nobullet inline"><a href="/news/auto/">Транспорт</a></li>
				<li class="nobullet inline"><a href="/news/company/">Новости компаний</a></li>
				<li class="nobullet inline"><a href="/news/socium/">Другая жизнь</a></li>
				<li class="nobullet inline"><a href="/news/article/">Статьи</a></li>
			</ul>
	</nav>
	</div>
</header>

<div class="row" id="wrapper">
	
	<div class="grid_8">
		<div  id="search" class="grid_8 row hide-phone">
		<div id="date" class="grid_2"><!-- Пятница, 28&nbsp;Августа&nbsp;2020 --></div>
		<div class="grid_6 text-right">
			<form method="get" class="search" action="/search/">
				<div class="wrapper-block"><input name="text" class="textbox" type="text" value="" />
  				<input class="submit btn-form" value="Поиск" type="submit" /></div>
			</form>
		</div>
		</div>
	
	<div class="grid_8 row">
	<aside class="grid_2 sidebar hide-phone">
		<div class="news_list ekb">
			<h2 class="block_title text-center">Новости Екатеринбурга</h2>
			
<h3><a href="/news/beauty/2020-05-22/id_64963.html">Профилактика вирусных инфекций в рекомендациях кандидата медицинских наук, врача-инфекциониста высшей категории Щинова Андрея Ивановича</a>&nbsp;<sub><a href="/news/beauty/2020-05-22/id_64963.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/auto/2020-05-21/id_62914.html">Пошив и продажа авточехлов для иномарок и отечественных автомобилей оптом и в розницу в Екатеринбурге</a>&nbsp;<sub><a href="/news/auto/2020-05-21/id_62914.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/line/2019-12-30/id_64954.html">Коллектив салонов оптики «ОЧКИ КУДЕЛИНОЙ» поздравляет горожан с наступающим 2020 годом!</a>&nbsp;<sub><a href="/news/line/2019-12-30/id_64954.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/sport/2019-11-09/id_64950.html">Благотворительный фонд Антона Шипулина и Благотворительный фонд «Общества Малышева-73 «Добрые дела» откроют еще 4 спортивных объекта</a>&nbsp;<sub><a href="/news/sport/2019-11-09/id_64950.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/education/2019-11-07/id_64949.html">В Екатеринбурге пройдет шестой Свердловский химический турнир</a>&nbsp;<sub><a href="/news/education/2019-11-07/id_64949.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/sport/2019-11-07/id_64948.html">Антон Шипулин и Игорь Заводовский откроют в Екатеринбурге три новые площадки для занятий по воркауту</a>&nbsp;<sub><a href="/news/sport/2019-11-07/id_64948.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/line/2019-11-07/id_64947.html">Межрегиональная общественная организация «Аистенок» проводит благотворительную акцию «Чашка тепла»</a>&nbsp;<sub><a href="/news/line/2019-11-07/id_64947.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/line/2019-10-02/id_64944.html">Екатеринбургский центр занятости приглашает граждан предпенсионного и пенсионного возраста на ярмарку вакансий</a>&nbsp;<sub><a href="/news/line/2019-10-02/id_64944.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/line/2019-09-09/id_64943.html">Благотворительный фонд Антона Шипулина предлагает студентам возможность воплотить свои идеи и получить финансовую поддержку</a>&nbsp;<sub><a href="/news/line/2019-09-09/id_64943.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/auto/2019-08-07/id_64937.html">В Екатеринбурге перекроют движение по Красноармейской, Калинина, Шейнкмана, Попова, Мичурина и Тверитина</a>&nbsp;<sub><a href="/news/auto/2019-08-07/id_64937.html#comments" title="Комментарии ()"></a></sub></h3>


			<div class="more"><a href="/news/eburg">Все новости Екатеринбурга&hellip;</a></div>
		</div>
		<div class="banner" id="banner_42"><a href="/banners/42/" title="" target="_blank" ><img src="/images/banners/61255_original.jpg" /></a></div>
<!--noindex-->
		<div class="text-center">

<!--LiveInternet counter--><script type="text/javascript"><!--
document.write("<a href='//www.liveinternet.ru/click' "+
"target=_blank><img src='//counter.yadro.ru/hit?t44.1;r"+
escape(document.referrer)+((typeof(screen)=="undefined")?"":
";s"+screen.width+"*"+screen.height+"*"+(screen.colorDepth?
screen.colorDepth:screen.pixelDepth))+";u"+escape(document.URL)+
";"+Math.random()+
"' alt='' title='LiveInternet' "+
"border='0' width='31' height='31'><\/a>")
//--></script><!--/LiveInternet-->
		</div>
<!--/noindex-->
	</aside>
	<div class="grid_6" id="content">
		<div class="grid_6 theme_day_block">
			
<article class="theme_day clear-left">
	<header><h1><a href="/news/health/2020-05-24/id_64821.html">Какие услуги можно получить по полису ОМС (обязательного медицинского страхования) бесплатно</a></h1></header>

<a href="/news/health/2020-05-24/id_64821.html"><figure><img src="/images/foto/72787_preview.jpg" class="img-left" alt="Какие услуги можно получить по полису ОМС (обязательного медицинского страхования) бесплатно" title="Какие услуги можно получить по полису ОМС (обязательного медицинского страхования) бесплатно" /></figure></a>

            <div class="news_desc"><a href="/news/health/2020-05-24/id_64821.html">Топ-10 бесплатных медицинских услуг, за которые могут неправомерно потребовать оплату в медицинских организациях, работающих в системе ОМС</a></div>
            <footer class="comments align-right">
                24.05.2020&nbsp;16:14 <br /><!--
                <a href="/news/health/2020-05-24/id_64821.html">Читать далее &raquo;</a>
				<a href="/news/health/2020-05-24/id_64821.html#comments">Комментарии ()</a>-->
			</footer>

</article>

		</div>
		<aside class="actual grid_6">
		<div class="grid_6 row">
			<div class="grid_6 row">
<div class="grid_3">
<article class="theme_day clear-left">
	<header><h1><a href="/news/beauty/2020-05-22/id_64963.html">Профилактика вирусных инфекций в рекомендациях кандидата медицинских наук, врача-инфекциониста высшей категории Щинова Андрея Ивановича</a></h1></header>

<a href="/news/beauty/2020-05-22/id_64963.html"><figure><img src="/images/foto/73048_square.jpg" class="img-left" alt="" title="" /></figure></a>
            <div class="news_desc"><a href="/news/beauty/2020-05-22/id_64963.html">Предлагаются оптимальные способы усиления жизненных сил организма, необходимых для успешной борьбы с этими коварными инфекциями.</a></div>
</article>
</div>

<div class="grid_3">
<article class="theme_day clear-left">
	<header><h1><a href="/news/auto/2020-05-21/id_62914.html">Пошив и продажа авточехлов для иномарок и отечественных автомобилей оптом и в розницу в Екатеринбурге</a></h1></header>

<a href="/news/auto/2020-05-21/id_62914.html"><figure><img src="/images/foto/70328_square.jpg" class="img-left" alt="" title="" /></figure></a>
            <div class="news_desc"><a href="/news/auto/2020-05-21/id_62914.html">Мы работаем для Вас уже 14 лет</a></div>
</article>
</div>
</div><div class="grid_6 row">
<div class="grid_3">
<article class="theme_day clear-left">
	<header><h1><a href="/news/line/2019-12-30/id_64954.html">Коллектив салонов оптики «ОЧКИ КУДЕЛИНОЙ» поздравляет горожан с наступающим 2020 годом!</a></h1></header>

<a href="/news/line/2019-12-30/id_64954.html"><figure><img src="/images/foto/73031_square.jpg" class="img-left" alt="" title="" /></figure></a>
            <div class="news_desc"><a href="/news/line/2019-12-30/id_64954.html">Опытные врачи офтальмологи более 15 лет помогают людям увидеть мир во всем его многообразии</a></div>
</article>
</div>

<div class="grid_3">
<article class="theme_day clear-left">
	<header><h1><a href="/news/sport/2019-11-09/id_64950.html">Благотворительный фонд Антона Шипулина и Благотворительный фонд «Общества Малышева-73 «Добрые дела» откроют еще 4 спортивных объекта</a></h1></header>

<a href="/news/sport/2019-11-09/id_64950.html"><figure><img src="/images/foto/73024_square.jpg" class="img-left" alt="" title="" /></figure></a>
            <div class="news_desc"><a href="/news/sport/2019-11-09/id_64950.html">В понедельник, 11 ноября в рамках программы «Территория спортивных побед»</a></div>
</article>
</div>
</div>
		</div>
		</aside>
		
		<aside class="actual grid_6">
		<div class="grid_6 row">
			<div class="grid_6 row">
<div class="grid_3">
<article class="theme_day clear-left">
	<header><h1><a href="/news/health/2020-05-24/id_64821.html">Какие услуги можно получить по полису ОМС (обязательного медицинского страхования) бесплатно</a></h1></header>

<a href="/news/health/2020-05-24/id_64821.html"><figure><img src="/images/foto/72787_square.jpg" class="img-left" alt="" title="" /></figure></a>
            <div class="news_desc"><a href="/news/health/2020-05-24/id_64821.html">Топ-10 бесплатных медицинских услуг, за которые могут неправомерно потребовать оплату в медицинских организациях, работающих в системе ОМС</a></div>
</article>
</div>

<div class="grid_3">
<article class="theme_day clear-left">
	<header><h1><a href="/news/company/2019-05-20/id_64913.html">Телеканал «Дождь» запускает новую линейку программ и объявляет специальные цены на подписку для регионов России</a></h1></header>

<a href="/news/company/2019-05-20/id_64913.html"><figure><img src="/images/foto/72949_square.jpg" class="img-left" alt="" title="" /></figure></a>
            <div class="news_desc"><a href="/news/company/2019-05-20/id_64913.html">В ближайшее время запланированы несколько премьер</a></div>
</article>
</div>
</div><div class="grid_6 row">
<div class="grid_3">
<article class="theme_day clear-left">
	<header><h1><a href="/news/polit/2019-04-30/id_64915.html">На Урале рыболовы поймали и отпустили почти 7,5 тысячи форелей</a></h1></header>

<a href="/news/polit/2019-04-30/id_64915.html"><figure><img src="/images/foto/72951_square.jpg" class="img-left" alt="" title="" /></figure></a>
            <div class="news_desc"><a href="/news/polit/2019-04-30/id_64915.html">За время соревнований было поймано и отпущено 7449 «хвостов». Фото: организаторы турнира, РК "Азарт"</a></div>
</article>
</div>

<div class="grid_3">
<article class="theme_day clear-left">
	<header><h1><a href="/news/auto/2019-04-10/id_64887.html">ГИБДД предложила увеличить минимальный штраф за превышение скорости до 3000 рублей и отменить скидку</a></h1></header>

<a href="/news/auto/2019-04-10/id_64887.html"><figure><img src="/images/foto/72898_square.jpg" class="img-left" alt="" title="" /></figure></a>
            <div class="news_desc"><a href="/news/auto/2019-04-10/id_64887.html">за быструю оплату штрафа</a></div>
</article>
</div>
</div><div class="grid_6 row">
<div class="grid_3">
<article class="theme_day clear-left">
	<header><h1><a href="/news/education/2019-04-04/id_64886.html">РЭУ им. Г.В.Плеханова по результатам аудита получил категорию «5 звезд» в рейтинге QS STARS</a></h1></header>

<a href="/news/education/2019-04-04/id_64886.html"><figure><img src="/images/foto/72896_square.jpg" class="img-left" alt="" title="" /></figure></a>
            <div class="news_desc"><a href="/news/education/2019-04-04/id_64886.html">став 47 вузом в мире, который имеет эту категорию</a></div>
</article>
</div>

<div class="grid_3">
<article class="theme_day clear-left">
	<header><h1><a href="/news/line/2019-03-20/id_64878.html">В Тюмени и Тобольске пройдёт отборочный этап Всероссийского конкурса детского и юношеского творчества «Земля талантов»</a></h1></header>

<a href="/news/line/2019-03-20/id_64878.html"><figure><img src="/images/foto/72879_square.jpg" class="img-left" alt="" title="" /></figure></a>
            <div class="news_desc"><a href="/news/line/2019-03-20/id_64878.html">В конкурсе могут принять участие дети и подростки следующих возрастных категорий: с 7 до 10 лет, с 11 до 14, с 15 до 17 (включительно)</a></div>
</article>
</div>
</div>
		</div>
		</aside>

		<aside class="grid_6 row education hide-phone">

		<h2 class="block_title"><span>Образование</span></h2>
			
<article class="theme_day clear-left">
	<a href="/news/education/2019-12-28/id_64953.html"><figure><img src="/images/foto/73030_preview.jpg" class="img-left" alt="" title="" /></figure></a>
	<header><h3><a href="/news/education/2019-12-28/id_64953.html">Пять главных принципов оформления помещения в частном детском саду</a></h3></header>
            <div><a href="/news/education/2019-12-28/id_64953.html"></a></div>
</article>
            

<article class="theme_day clear-left">
	<a href="/news/education/2019-11-07/id_64949.html"><figure><img src="/images/foto/73023_preview.jpg" class="img-left" alt="" title="" /></figure></a>
	<header><h3><a href="/news/education/2019-11-07/id_64949.html">В Екатеринбурге пройдет шестой Свердловский химический турнир</a></h3></header>
            <div><a href="/news/education/2019-11-07/id_64949.html"></a></div>
</article>
            

<article class="theme_day clear-left">
	<a href="/news/education/2019-07-08/id_64929.html"><figure><img src="/images/foto/72969_preview.jpg" class="img-left" alt="" title="" /></figure></a>
	<header><h3><a href="/news/education/2019-07-08/id_64929.html">Иностранный язык с самого раннего возраста в Центре "Полиглотики"!</a></h3></header>
            <div><a href="/news/education/2019-07-08/id_64929.html"></a></div>
</article>
            

			<div class="more"><a href="/news/education/">Читать далее&hellip;</a></div>

			
		</aside>
		<aside class="grid_6 row article hide-phone">
		<h2 class="block_title"><span>Статьи</span></h2>
			
<article class="theme_day clear-left">
	<a href="/news/home_family/2020-07-22/id_64974.html"><figure><img src="/images/foto/73062_preview.jpg" class="img-left" alt="" title="" /></figure></a>
	<header><h3><a href="/news/home_family/2020-07-22/id_64974.html">Дизайн проекты домов</a></h3></header>
            <div><a href="/news/home_family/2020-07-22/id_64974.html"></a></div>
</article>
            

<article class="theme_day clear-left">
	<a href="/news/animal/2020-07-15/id_64973.html"><figure><img src="/images/foto/73063_preview.jpg" class="img-left" alt="" title="" /></figure></a>
	<header><h3><a href="/news/animal/2020-07-15/id_64973.html">Промышленные корма для собак и кошек на рынке: мнения профессионалов.</a></h3></header>
            <div><a href="/news/animal/2020-07-15/id_64973.html"></a></div>
</article>
            

<article class="theme_day clear-left">
	<a href="/news/business_interests/2020-07-07/id_64972.html"><figure><img src="/images/foto/73060_preview.jpg" class="img-left" alt="" title="" /></figure></a>
	<header><h3><a href="/news/business_interests/2020-07-07/id_64972.html">Проверка задолженности в ФНС.</a></h3></header>
            <div><a href="/news/business_interests/2020-07-07/id_64972.html"></a></div>
</article>
            

			<div class="more"><a href="/news/article/">Читать далее&hellip;</a></div>
		</aside>
		
	<aside id="tags" class="grid_6 hide-phone">
		<h2 class="block_title"><a href="/news/tag/">Облако тегов</a></h2>
		<p class="tags"><a href="/news/tag/%ED%EE%E2%EE%F1%F2%E8+%CA%F0%FB%EC%E0" class="weight1">новости&nbsp;Крыма</a> <a href="/news/tag/%ED" class="weight1">н</a> <a href="/news/tag/%ED%EE%E2%EE%F1%F2%E8+%D5%E0%F0%FC%EA%EE%E2%E0" class="weight1">новости&nbsp;Харькова</a> <a href="/news/tag/%ED%EE%E2%EE%F1%F2%E8+%E4%EE%EC+2+2013" class="weight1">новости&nbsp;дом&nbsp;2&nbsp;2013</a> <a href="/news/tag/%E0%E2%E0%F0%E8%FF+%E2+%C5%EA%E0%F2%E5%F0%E8%ED%E1%F3%F0%E3%E5" class="weight1">авария&nbsp;в&nbsp;Екатеринбурге</a> <a href="/news/tag/%F4%F3%F2%E1%EE%EB+%F7%E5%EC%EF%E8%EE%ED%E0%F2+%D0%EE%F1%F1%E8%E8" class="weight1">футбол&nbsp;чемпионат&nbsp;России</a> <a href="/news/tag/%C1%E8%E0%F2%EB%EE%ED" class="weight1">Биатлон</a> <a href="/news/tag/%ED%EE%E2%EE%F1%F2%E8+%E4%EE%EC+2+%F1%E5%E3%EE%E4%ED%FF" class="weight1">новости&nbsp;дом&nbsp;2&nbsp;сегодня</a> <a href="/news/tag/%E1%E5%F1%EF%EE%F0%FF%E4%EA%E8+%E2+%C4%EE%ED%E5%F6%EA%E5" class="weight1">беспорядки&nbsp;в&nbsp;Донецке</a> <a href="/news/tag/%ED%EE%E2%EE%F1%F2%E8+%D1%E5%E2%E0%F1%F2%EE%EF%EE%EB%FF" class="weight1">новости&nbsp;Севастополя</a> <a href="/news/tag/%F1%E0%ED%EA%F6%E8%E8" class="weight1">санкции</a> <a href="/news/tag/%CF%EE%E6%E0%F0" class="weight1">Пожар</a> <a href="/news/tag/%D1%D8%C0" class="weight1">США</a> <a href="/news/tag/%ED%EE%E2%EE%F1%F2%E8+%EA%E8%E5%E2%E0" class="weight4">новости&nbsp;киева</a> <a href="/news/tag/%ED%EE%E2%EE%F1%F2%E8+%C4%E0%E3%E5%F1%F2%E0%ED%E0" class="weight1">новости&nbsp;Дагестана</a> <a href="/news/tag/%ED%EE%E2%EE%F1%F2%E8+%C4%EE%ED%E5%F6%EA%E0" class="weight2">новости&nbsp;Донецка</a> <a href="/news/tag/%D1%E2%E5%F0%E4%EB%EE%E2%F1%EA%E0%FF+%EE%E1%EB%E0%F1%F2%FC" class="weight2">Свердловская&nbsp;область</a> <a href="/news/tag/%EA%E8%E5%E2+%ED%EE%E2%EE%F1%F2%E8" class="weight1">киев&nbsp;новости</a> <a href="/news/tag/%EE%EF%E5%F0%E0%F6%E8%FF+%E2+%D1%EB%E0%E2%FF%ED%F1%EA%E5" class="weight1">операция&nbsp;в&nbsp;Славянске</a> <a href="/news/tag/%ED%EE%E2%EE%F1%F2%E8+%CB%F3%E3%E0%ED%F1%EA%E0+%F1%E5%E3%EE%E4%ED%FF" class="weight2">новости&nbsp;Луганска&nbsp;сегодня</a> <a href="/news/tag/%ED%EE%E2%EE%F1%F2%E8+%CA%E8%E5%E2%E0+%F1%E5%E3%EE%E4%ED%FF" class="weight1">новости&nbsp;Киева&nbsp;сегодня</a> <a href="/news/tag/%D4%F3%F2%E1%EE%EB" class="weight3">Футбол</a> <a href="/news/tag/%ED%EE%E2%EE%F1%F2%E8+%CA%F0%FB%EC%E0+%F1%E5%E3%EE%E4%ED%FF" class="weight1">новости&nbsp;Крыма&nbsp;сегодня</a> <a href="/news/tag/%ED%EE%E2%EE%F1%F2%E8+%E1%E8%E0%F2%EB%EE%ED%E0" class="weight1">новости&nbsp;биатлона</a> <a href="/news/tag/%ED%EE%E2%EE%F1%F2%E8+%E4%EE%EC+2+%ED%E0+%F1%E5%E3%EE%E4%ED%FF" class="weight2">новости&nbsp;дом&nbsp;2&nbsp;на&nbsp;сегодня</a> <a href="/news/tag/%C4%D2%CF+%E2+%C5%EA%E0%F2%E5%F0%E8%ED%E1%F3%F0%E3%E5" class="weight1">ДТП&nbsp;в&nbsp;Екатеринбурге</a> <a href="/news/tag/%C4%D2%CF" class="weight4">ДТП</a> <a href="/news/tag/%ED%EE%E2%EE%F1%F2%E8+%D1%EB%E0%E2%FF%ED%F1%EA%E0" class="weight1">новости&nbsp;Славянска</a> <a href="/news/tag/%C0%EB%E5%EA%F1%E0%ED%E4%F0+%CC%E8%F8%E0%F0%E8%ED" class="weight1">Александр&nbsp;Мишарин</a> <a href="/news/tag/%E4%EE%EC+2+%ED%EE%E2%EE%F1%F2%E8" class="weight1">дом&nbsp;2&nbsp;новости</a> <a href="/news/tag/%EF%F0%EE%E8%F8%E5%F1%F2%E2%E8%FF+%E2+%EC%EE%F1%EA%E2%E5" class="weight1">проишествия&nbsp;в&nbsp;москве</a> <a href="/news/tag/%F1%EF%EB%E5%F2%ED%E8+%E4%EE%EC+2" class="weight2">сплетни&nbsp;дом&nbsp;2</a> <a href="/news/tag/%C5%E2%E3%E5%ED%E8%E9+%CA%F3%E9%E2%E0%F8%E5%E2" class="weight1">Евгений&nbsp;Куйвашев</a> <a href="/news/tag/%ED%EE%E2%EE%F1%F2%E8+%CE%E4%E5%F1%F1%FB" class="weight1">новости&nbsp;Одессы</a> <a href="/news/tag/%ED%EE%E2%EE%F1%F2%E8+%E4%EE%EC+2+2012" class="weight1">новости&nbsp;дом&nbsp;2&nbsp;2012</a> <a href="/news/tag/%ED%EE%E2%EE%F1%F2%E8+%E4%EE%EC+2" class="weight2">новости&nbsp;дом&nbsp;2</a> <a href="/news/tag/%C2%EB%E0%E4%E8%EC%E8%F0+%CF%F3%F2%E8%ED" class="weight1">Владимир&nbsp;Путин</a> <a href="/news/tag/%EF%F0%EE%E8%F1%F8%E5%F1%F2%E2%E8%FF+%E2+%CC%EE%F1%EA%E2%E5" class="weight1">происшествия&nbsp;в&nbsp;Москве</a> <a href="/news/tag/%F7%F2%EE+%EF%F0%EE%E8%F1%F5%EE%E4%E8%F2+%E2+%D1%EB%E0%E2%FF%ED%F1%EA%E5" class="weight1">что&nbsp;происходит&nbsp;в&nbsp;Славянске</a> <a href="/news/tag/%E4%EE%ED%E5%F6%EA+%ED%EE%E2%EE%F1%F2%E8+%F1%E5%E3%EE%E4%ED%FF" class="weight2">донецк&nbsp;новости&nbsp;сегодня</a> <a href="/news/tag/%EF%EE%E3%EE%E4%E0+%E2+%CC%EE%F1%EA%E2%E5" class="weight1">погода&nbsp;в&nbsp;Москве</a> <a href="/news/tag/%E0%E2%E0%F0%E8%FF" class="weight3">авария</a> <a href="/news/tag/%ED%EE%E2%EE%F1%F2%E8+%EC%EE%F1%EA%E2%FB" class="weight2">новости&nbsp;москвы</a> <a href="/news/tag/%F1%EE%F1%F2%E0%E2+%F1%E1%EE%F0%ED%EE%E9+%D0%EE%F1%F1%E8%E8" class="weight1">состав&nbsp;сборной&nbsp;России</a> <a href="/news/tag/%ED%EE%E2%EE%F1%F2%E8+%D3%EA%F0%E0%E8%ED%FB" class="weight5">новости&nbsp;Украины</a> <a href="/news/tag/%ED%EE%E2%EE%F1%F2%E8+%F4%F3%F2%E1%EE%EB%E0" class="weight3">новости&nbsp;футбола</a> <a href="/news/tag/%F1%EB%F3%F5%E8+%ED%E0+%E4%EE%EC+2" class="weight2">слухи&nbsp;на&nbsp;дом&nbsp;2</a> <a href="/news/tag/%F1%E0%ED%EA%F6%E8%E8+%EF%F0%EE%F2%E8%E2+%D0%EE%F1%F1%E8%E8" class="weight1">санкции&nbsp;против&nbsp;России</a> <a href="/news/tag/%ED%EE%E2%EE%F1%F2%E8+%F3%EA%F0%E0%E8%ED%FB+%F1%E5%E3%EE%E4%ED%FF" class="weight4">новости&nbsp;украины&nbsp;сегодня</a> <a href="/news/tag/%ED%EE%E2%EE%F1%F2%E8+%E4%EE%EC+2+%F1%E2%E5%E6%E8%E5" class="weight2">новости&nbsp;дом&nbsp;2&nbsp;свежие</a> <a href="/news/tag/%E4%EE%EC+2+%F1%E2%E5%E6%E8%E5+%ED%EE%E2%EE%F1%F2%E8" class="weight2">дом&nbsp;2&nbsp;свежие&nbsp;новости</a> <a href="/news/tag/%E1%E5%F1%EF%EE%F0%FF%E4%EA%E8+%E2+%CA%E8%E5%E2%E5" class="weight1">беспорядки&nbsp;в&nbsp;Киеве</a> <a href="/news/tag/%ED%EE%E2%EE%F1%F2%E8+%CF%E5%F2%E5%F0%E1%F3%F0%E3%E0" class="weight1">новости&nbsp;Петербурга</a> <a href="/news/tag/%F3%E1%E8%E9%F1%F2%E2%EE" class="weight1">убийство</a> <a href="/news/tag/%ED%EE%E2%EE%F1%F2%E8+%CA%F0%E0%EC%E0%F2%EE%F0%F1%EA%E0" class="weight1">новости&nbsp;Краматорска</a> <a href="/news/tag/%F3%EA%F0%E0%E8%ED%E0+%ED%EE%E2%EE%F1%F2%E8+%F1%E5%E3%EE%E4%ED%FF" class="weight4">украина&nbsp;новости&nbsp;сегодня</a> <a href="/news/tag/%F4%F3%F2%E1%EE%EB+%D0%EE%F1%F1%E8%E8" class="weight2">футбол&nbsp;России</a> <a href="/news/tag/%F4%F3%F2%E1%EE%EB+%F1%E5%E3%EE%E4%ED%FF" class="weight2">футбол&nbsp;сегодня</a> <a href="/news/tag/%ED%EE%E2%EE%F1%F2%E8+%CB%F3%E3%E0%ED%F1%EA%E0" class="weight1">новости&nbsp;Луганска</a> <a href="/news/tag/%ED%EE%E2%EE%F1%F2%E8+%E5%EA%E0%F2%E5%F0%E8%ED%E1%F3%F0%E3%E0" class="weight7">новости&nbsp;екатеринбурга</a> </p>
		<p align="center"><a href="/">Главная</a> | <a href="http://0-50.ru/news/stock">Карта сайта</a> | <a href="/rss/">RSS Feed</a></p>
	</aside>
	</div><!-- end #content -->
	</div>
		
		
	</div>
	<aside class="grid_4">
		
	</aside>
	<aside class="grid_2 news_list sidebar">
	<div class="banner" id="banner_56"><p><a href="https://www.kartridgy.ru/" target="_blank"><img src="/images/bannerkr.jpg" width="200" alt="" /></a></p></div>
	<h2 class="block_title text-center">Россия и мир</h2>
	
<h3><a href="/news/health/2020-05-24/id_64821.html">Какие услуги можно получить по полису ОМС (обязательного медицинского страхования) бесплатно</a>&nbsp;<sub><a href="/news/health/2020-05-24/id_64821.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/company/2019-05-20/id_64913.html">Телеканал «Дождь» запускает новую линейку программ и объявляет специальные цены на подписку для регионов России</a>&nbsp;<sub><a href="/news/company/2019-05-20/id_64913.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/polit/2019-04-30/id_64915.html">На Урале рыболовы поймали и отпустили почти 7,5 тысячи форелей</a>&nbsp;<sub><a href="/news/polit/2019-04-30/id_64915.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/auto/2019-04-10/id_64887.html">ГИБДД предложила увеличить минимальный штраф за превышение скорости до 3000 рублей и отменить скидку</a>&nbsp;<sub><a href="/news/auto/2019-04-10/id_64887.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/education/2019-04-04/id_64886.html">РЭУ им. Г.В.Плеханова по результатам аудита получил категорию «5 звезд» в рейтинге QS STARS</a>&nbsp;<sub><a href="/news/education/2019-04-04/id_64886.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/line/2019-03-20/id_64878.html">В Тюмени и Тобольске пройдёт отборочный этап Всероссийского конкурса детского и юношеского творчества «Земля талантов»</a>&nbsp;<sub><a href="/news/line/2019-03-20/id_64878.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/line/2019-03-06/id_64869.html">Подарки на поздравление с 8 марта каждый третий россиянин ищет в интернете</a>&nbsp;<sub><a href="/news/line/2019-03-06/id_64869.html#comments" title="Комментарии ()"></a></sub></h3>

<div id="unit_85265"><a href="http://smi2.net/">Новости СМИ2</a></div>
<script type="text/javascript" charset="utf-8">
  (function() {
    var sc = document.createElement('script'); sc.type = 'text/javascript'; sc.async = true;
    sc.src = '//news.smi2.ru/data/js/85265.js'; sc.charset = 'utf-8';
    var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(sc, s);
  }());
</script>
<h3><a href="/news/socium/2019-02-17/id_64860.html">Жена Марата Башарова Елизавета Шевыркова подала на развод</a>&nbsp;<sub><a href="/news/socium/2019-02-17/id_64860.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/incident/2019-01-08/id_64841.html">В Воронеже во дворе дома №77 на Калачеевской пьяный мужчина сжег Audi A8</a>&nbsp;<sub><a href="/news/incident/2019-01-08/id_64841.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/socium/2018-12-17/id_64830.html">Филиппинка Катриона Грэй победила в конкурсе красоты «Мисс Вселенная 2018», Юлия Полячихина не вышла даже в полуфинал</a>&nbsp;<sub><a href="/news/socium/2018-12-17/id_64830.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/line/2018-12-17/id_64829.html">В Торжке (Тверская область) парень получил срок за пост «ВКонтакте»</a>&nbsp;<sub><a href="/news/line/2018-12-17/id_64829.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/incident/2018-12-16/id_64827.html">Ехавший в Рязань мужчина на трассе А-108 спас подорвавшегося участкового ОМВД «Марьина Роща»</a>&nbsp;<sub><a href="/news/incident/2018-12-16/id_64827.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/incident/2018-12-16/id_64826.html">Взрыв метана на полигоне ТБО у Конаково под Тверью привел к пожару, который вряд ли ликвидируют в ближайшие дни</a>&nbsp;<sub><a href="/news/incident/2018-12-16/id_64826.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/line/2018-11-01/id_64784.html">Объявлен Грантовый конкурс музейных инклюзивных программ 2018</a>&nbsp;<sub><a href="/news/line/2018-11-01/id_64784.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/careers/2018-11-01/id_64783.html">Как правильно оформить страховку?</a>&nbsp;<sub><a href="/news/careers/2018-11-01/id_64783.html#comments" title="Комментарии ()"></a></sub></h3>


	<div class="more"><a href="/news/">Все новости&hellip;</a></div>
	
	
	
	
	
	</aside>
	<aside class="grid_2 sidebar">
			
			<div class="news_list">
			<h2 class="block_title text-center">Актуально</h2>
			
<h3><a href="/news/health/2020-05-24/id_64821.html">Какие услуги можно получить по полису ОМС (обязательного медицинского страхования) бесплатно</a>&nbsp;<sub><a href="/news/health/2020-05-24/id_64821.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/beauty/2020-05-22/id_64963.html">Профилактика вирусных инфекций в рекомендациях кандидата медицинских наук, врача-инфекциониста высшей категории Щинова Андрея Ивановича</a>&nbsp;<sub><a href="/news/beauty/2020-05-22/id_64963.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/auto/2020-05-21/id_62914.html">Пошив и продажа авточехлов для иномарок и отечественных автомобилей оптом и в розницу в Екатеринбурге</a>&nbsp;<sub><a href="/news/auto/2020-05-21/id_62914.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/company/2019-05-15/id_64908.html">Екатеринбургский центр занятости предлагает горожанам записаться на консультацию на сайте Департамента по труду и занятости Свердловской области</a>&nbsp;<sub><a href="/news/company/2019-05-15/id_64908.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/company/2019-04-17/id_64891.html">Екатеринбургский центр занятости приглашает граждан предпенсионного возраста пройти бесплатное обучение</a>&nbsp;<sub><a href="/news/company/2019-04-17/id_64891.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/cooking/2019-03-05/id_64867.html">Что приготовить на ужин быстро и вкусно: салат из куриной грудки рецепт с фото своими руками или на заказ</a>&nbsp;<sub><a href="/news/cooking/2019-03-05/id_64867.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/company/2018-12-27/id_64840.html">Межрегиональная общественная организация «Аистенок» проводит благотворительную акцию «Мама, останься!»</a>&nbsp;<sub><a href="/news/company/2018-12-27/id_64840.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/health/2018-12-09/id_64818.html">Лечение синдрома хронической усталости с помощью дыхательных упражнений</a>&nbsp;<sub><a href="/news/health/2018-12-09/id_64818.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/cooking/2018-04-26/id_64440.html">Торт из морковки простой рецепт с фото своими руками или на заказ</a>&nbsp;<sub><a href="/news/cooking/2018-04-26/id_64440.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/company/2017-05-09/id_62916.html">Имущественные споры граждан при разводе (расторжении брака) в мировом суде: «Эксперт-Ком» Екатеринбург</a>&nbsp;<sub><a href="/news/company/2017-05-09/id_62916.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/cooking/2017-03-31/id_62729.html">Что приготовить на ужин быстро и вкусно: куриные маффины с грибами рецепт с фото пошагово</a>&nbsp;<sub><a href="/news/cooking/2017-03-31/id_62729.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/cooking/2017-03-20/id_62656.html">Что приготовить на ужин быстро и вкусно: открытый пирог с семгой и брокколи рецепт с фото</a>&nbsp;<sub><a href="/news/cooking/2017-03-20/id_62656.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/cooking/2017-03-19/id_62644.html">Что приготовить на ужин быстро и вкусно: ризотто с грибами и курицей рецепт с фото пошагово</a>&nbsp;<sub><a href="/news/cooking/2017-03-19/id_62644.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/cooking/2017-03-05/id_62551.html">Что приготовить на ужин из мяса быстро и вкусно: рулетики с грибами и вишней рецепт с фото пошагово</a>&nbsp;<sub><a href="/news/cooking/2017-03-05/id_62551.html#comments" title="Комментарии ()"></a></sub></h3>


<h3><a href="/news/cooking/2017-02-20/id_62489.html">Что приготовить на ужин 23 февраля и 8 марта быстро и вкусно: грудинка на кости рецепт с фото пошагово</a>&nbsp;<sub><a href="/news/cooking/2017-02-20/id_62489.html#comments" title="Комментарии ()"></a></sub></h3>


			<div class="more"><a href="/news/actual">Все актуальные сюжеты&hellip;</a></div>
			</div>
	
	<div class="banner" id="banner_51"><a href="/banners/51/" title="" target="_blank" ><img src="/images/banners/69250_original.jpg" /></a></div>
	
	
	</aside>
</div> <!-- end #wrapper -->

<footer class="row">
	<div class="grid_12">
			<p style="font-size:11px; line-height: 13px;">Сетевое издание Служба новостей 050&nbsp; <br>зарегистрировано в Федеральной службе по надзору в сфере связи,<br>
информационных технологий и массовых коммуникаций (Роскомнадзор) 25 апреля 2017г.<br>
Свидетельство о регистрации&nbsp;ЭЛ № ФС77-69503<br />
<!--Учредитель&nbsp;общество с ограниченной ответственностью &laquo;Горсправка-ИНФО&raquo;</p> -->

<p style="font-size:11px; line-height: 13px;">Главный редактор&nbsp;Давыдов А.В.<br />
Руководитель проекта&nbsp;Николаев В.П.<br />
Адрес электронной почты редакции&nbsp;<a href="mailto:1@0-50.ru">1@0-50.ru</a><br />
Телефон редакции&nbsp;+7 (908)&nbsp;914-27-00<br />
Настоящий ресурс содержит материалы&nbsp;18+</p>

<p>© Горсправка 2009-2019</p>





	</div>
</footer>
</body>
</html>`)

	expected := &Company{
		Title: "Новости 0-50.ru |",
		Email: "1@0-50.ru",
		Phone: 79089142700,
	}
	actual := &Company{}
	actual.digHTML(context.Background(), news050ruHTML, true, false, false)

	assert.Equal(t, expected, actual)
}
