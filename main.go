package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

func main() {

	// Загрузим страницы
	for url, path := range getInfo() {
		if err := download(url, path); err != nil {
			log.Fatalln(err)
		}
	}

	r()

}

func getInfo() map[string]string {
	return map[string]string{
		"https://dentalik.webflow.io":          "D:\\PROG\\GoProjects\\src\\WFDL\\out\\index.html",
		"https://dentalik.webflow.io/products": "D:\\PROG\\GoProjects\\src\\WFDL\\out\\products.html",
	}
}

func download(url string, path string) (err error) {

	var resp *http.Response
	var file *os.File

	resp, err = http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err = os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	io.Copy(file, resp.Body)

	return err
}

func r() {
	var str = `<!DOCTYPE html><!-- This site was created in Webflow. http://www.webflow.com --><!-- Last Published: Fri Mar 06 2020 05:03:22 GMT+0000 (Coordinated Universal Time) --><html data-wf-domain="dentalik.webflow.io" data-wf-page="5e4e4fa845b75d4de21305a8" data-wf-site="5e4e4fa75e5be328c3a7fad2" data-wf-status="1"><head><meta charset="utf-8"/><title>ООО Денталик</title><meta content="ООО Денталик" property="og:title"/><meta content="width=device-width, initial-scale=1" name="viewport"/><meta content="Webflow" name="generator"/><link href="https://uploads-ssl.webflow.com/5e4e4fa75e5be328c3a7fad2/css/dentalik.webflow.2c64c83c1.css" rel="stylesheet" type="text/css"/><script src="https://ajax.googleapis.com/ajax/libs/webfont/1.6.26/webfont.js" type="text/javascript"></script><script type="text/javascript">WebFont.load({  google: {    families: ["Open Sans:300,300italic,400,400italic,600,600italic,700,700italic,800,800italic"]  }});</script><!--[if lt IE 9]><script src="https://cdnjs.cloudflare.com/ajax/libs/html5shiv/3.7.3/html5shiv.min.js" type="text/javascript"></script><![endif]--><script type="text/javascript">!function(o,c){var n=c.documentElement,t=" w-mod-";n.className+=t+"js",("ontouchstart"in o||o.DocumentTouch&&c instanceof DocumentTouch)&&(n.className+=t+"touch")}(window,document);</script><link href="https://uploads-ssl.webflow.com/5e4e4fa75e5be328c3a7fad2/5e4e4fe82616d379f1a48544_5e4ba7853a237e3732c780a3_favicon.png" rel="shortcut icon" type="image/x-icon"/><link href="https://uploads-ssl.webflow.com/5e4e4fa75e5be328c3a7fad2/5e4e4fed47c62005d72f07e9_5e4ba82c1e21af3ab43adb89_favicon256x256.png" rel="apple-touch-icon"/><meta name="robots" content="noindex, nofollow">
<meta name="author" content="wtw"><!-- HEAD CODE --></head><body><div data-collapse="medium" data-animation="default" data-duration="400" id="menu" class="menu_main w-nav"><div class="container w-container"><a href="#" class="w-nav-brand"></a><nav role="navigation" class="nav-menu w-nav-menu"><a href="#" class="menu_main_item w-nav-link">Главная</a><div data-delay="0" class="w-dropdown"><div class="menu_main_item w-dropdown-toggle"><div class="icon-3 w-icon-dropdown-toggle"></div><div>продукция</div></div><nav class="menu_main_dropdown_list1 w-dropdown-list"><div data-delay="0" class="w-dropdown"><div class="menu_main_item w-dropdown-toggle"><div class="w-icon-dropdown-toggle"></div><div>Зубные пасты</div></div><nav class="menu_main_dropdown_list2 w-dropdown-list"><a href="/products/#toothpaste1" class="menu_main_item w-dropdown-link">Премиум-линия</a><a href="/products/#toothpaste2" class="menu_main_item w-dropdown-link">Нишевая линия</a><a href="/products/#toothpaste3" class="menu_main_item w-dropdown-link">Про-линия</a><a href="/products/#toothpaste4" class="menu_main_item w-dropdown-link">Детская линия</a></nav></div><div data-delay="0" class="w-dropdown"><div class="menu_main_item w-dropdown-toggle"><div class="w-icon-dropdown-toggle"></div><div>зубные щетки</div></div><nav class="menu_main_dropdown_list2 w-dropdown-list"><a href="/products/#toothbrush1" class="menu_main_item w-dropdown-link">Премиум-линия</a><a href="/products/#toothbrush2" class="menu_main_item w-dropdown-link">Про-линия</a><a href="/products/#toothbrush3" class="menu_main_item w-dropdown-link">Детская линия</a></nav></div><a href="/products/#rinsers" class="menu_main_item w-dropdown-link">ополаскиватели</a></nav></div><a href="/#about" class="menu_main_item w-nav-link">о нас</a><a href="/#contact" class="menu_main_item w-nav-link">контакты</a></nav><div class="menu-button w-nav-button"><div class="icon-2 w-icon-nav-menu"></div></div></div></div><div id="top" class="top"><div data-delay="6000" data-animation="cross" data-autoplay="1" data-easing="ease-in-quad" data-duration="3000" data-infinite="1" class="slider w-slider"><div class="w-slider-mask"><div class="w-slide"><div class="slide_bg _1" data-ix="slider"></div><div class="slide_content"><img src="https://uploads-ssl.webflow.com/5e4e4fa75e5be328c3a7fad2/5e4e50f5be890423cba2b49b_full_logo_white.png" alt="" class="slider_logo"/><img src="https://uploads-ssl.webflow.com/5e4e4fa75e5be328c3a7fad2/5e4e50f6eb37d1284ff7aa2f_separator.png" alt="" class="slider_separator"/><h1 class="slider_text">Исключительно<br/>растительные компоненты</h1></div></div><div class="w-slide"><div class="slide_bg _2" data-ix="slider"></div><div class="slide_content"><img src="https://uploads-ssl.webflow.com/5e4e4fa75e5be328c3a7fad2/5e4e50f5be890423cba2b49b_full_logo_white.png" alt="" class="slider_logo"/><img src="https://uploads-ssl.webflow.com/5e4e4fa75e5be328c3a7fad2/5e4e50f6eb37d1284ff7aa2f_separator.png" alt="" class="slider_separator"/><h1 class="slider_text">Не тестировалось<br/>на животных</h1></div></div><div class="w-slide" data-ix="slider"><div class="slide_bg _3" data-ix="slider"></div><div class="slide_content"><img src="https://uploads-ssl.webflow.com/5e4e4fa75e5be328c3a7fad2/5e4e50f5be890423cba2b49b_full_logo_white.png" alt="" class="slider_logo"/><img src="https://uploads-ssl.webflow.com/5e4e4fa75e5be328c3a7fad2/5e4e50f6eb37d1284ff7aa2f_separator.png" alt="" class="slider_separator"/><h1 class="slider_text">Идеально для Вас<br/>и Вашей семьи</h1></div></div></div><div class="left-arrow-2 w-slider-arrow-left"><div class="w-icon-slider-left"></div></div><div class="right-arrow-2 w-slider-arrow-right"><div class="w-icon-slider-right"></div></div><div class="w-slider-nav w-round"></div></div></div><div class="hits"><div class="w-container"><h2>ХИТЫ ПРОДАЖ</h2><div data-delay="5000" data-animation="cross" data-autoplay="1" data-easing="ease-in-quad" data-duration="500" data-infinite="1" class="slider-2 w-slider"><div class="w-slider-mask"><div class="hit_slide w-slide"><div class="hit"><img src="https://uploads-ssl.webflow.com/5e4e4fa75e5be328c3a7fad2/5e4e51c979f2006895a29028_BLACK.png" alt=""/><h5>ЗУБНАЯ ПАСТА EXTRA-WHITENING BLACK</h5></div></div><div class="hit_slide w-slide"><div class="hit"><img src="https://uploads-ssl.webflow.com/5e4e4fa75e5be328c3a7fad2/5e4e51c9be8904301ea2b866_GOLD_SWISS.png" alt=""/><h5>ЗУБНАЯ ПАСТА-ГЕЛЬ ADVANCED WHITENING GOLD SWITZERLAND</h5></div></div><div class="hit_slide w-slide"><div class="hit"><img src="https://uploads-ssl.webflow.com/5e4e4fa75e5be328c3a7fad2/5e4e51c947c6208bc42f0fb8_KIDS.png" alt=""/><h5>ЗУБНАЯ ПАСТА-ГЕЛЬ JUNIOR 6+ YEARS</h5></div></div><div class="hit_slide w-slide"><div class="hit"><img src="https://uploads-ssl.webflow.com/5e4e4fa75e5be328c3a7fad2/5e4e51c95e5be35cbaa807bd_PREGNANT.png" alt=""/><h5>ЗУБНАЯ ПАСТА-ГЕЛЬ PREGNANT LADY &amp; YOUNG MUM</h5></div></div><div class="hit_slide w-slide"><div class="hit"><img src="https://uploads-ssl.webflow.com/5e4e4fa75e5be328c3a7fad2/5e4e51c95e5be31020a807bc_VEGAN.png" alt=""/><h5>ЗУБНАЯ ПАСТА-ГЕЛЬ VEGAN NATURAL WITH VITAMIN B12</h5></div></div><div class="hit_slide w-slide"><div class="hit"><img src="https://uploads-ssl.webflow.com/5e4e4fa75e5be328c3a7fad2/5e4e51c9eb37d19835f7b0f5_PRO-WHITENING.png" alt=""/><h5>ЗУБНАЯ ПАСТА PRO-WHITENING</h5></div></div></div><div class="left-arrow w-slider-arrow-left"><div class="icon w-icon-slider-left"></div></div><div class="right-arrow w-slider-arrow-right"><div class="w-icon-slider-right"></div></div><div class="slide-nav w-slider-nav w-round"></div></div></div></div><section id="about" class="about"><div class="centered-container w-container"><h2>о нас</h2><p><strong>Medpack Swiss Group</strong> – это холдинг со штаб-квартирой в Швейцарии, 3-мя собственными производственными предприятиями, дочерними компаниями в 14-ти странах Европы, Ближнего Востока и Азии, а также дистрибьюторами в 44-х странах мира.</p><p>Ежедневно <strong>Medpack Swiss Group</strong> отгружает тысячи товаров из собственных логистических центров в ЕС (Амстердам, Рига, София), ОАЭ (Шарджа), Гонконг, Майами.</p><p>Наша компания ориентирована на долгосрочное сотрудничество, наилучшее соотношение «цена-качество» и полное удовлетворение потребностей наших клиентов.</p><div class="cards-grid-container"><div id="w-node-d99af7bd5676-e21305a8"><div class="cards-image-mask"><img src="https://uploads-ssl.webflow.com/5e4e4fa75e5be328c3a7fad2/5e4e50f65e5be34666a80389_manufacture.jpg" alt="" class="cards-image"/></div><h3>производство</h3><p>Продукция DENTISSIMO выпускается на самом современном заводе по высоким технологиям и GMP сертификату.<br/>‍<br/>В соответствии с правилами были установлены специальные надежные анти-дефлаграционные системы.<br/>Продукты DENTISSIMO производятся ежедневно в соответствии с детальными производственными процессами и сертифицированным контролем.</p></div><div id="w-node-d99af7bd5683-e21305a8"><div class="cards-image-mask"><img src="https://uploads-ssl.webflow.com/5e4e4fa75e5be328c3a7fad2/5e4e50f6be8904e3f5a2b49c_sunflower.jpg" alt="" class="cards-image"/></div><h3 id="w-node-d99af7bd5686-e21305a8">сделано с заботой</h3><p id="w-node-d99af7bd5688-e21305a8">В состав Dentissimo входят исключительно растительные компоненты, поэтому в составах зубных паст нет никаких ингредиентов животного происхождения.<br/><br/>Мы никогда не проводили тестирование на животных. Dentissimo пытается свести к минимуму воздействие на окружающую среду и, по возможности, делает вклад в определенные проекты на местном уровне и за рубежом.</p></div><div id="w-node-d99af7bd568d-e21305a8"><div class="cards-image-mask"><img src="https://uploads-ssl.webflow.com/5e4e4fa75e5be328c3a7fad2/5e4e50f50f3ef8eac561b215_lab.jpg" alt="" class="cards-image"/></div><h3 id="w-node-d99af7bd5690-e21305a8">что в составе?</h3><p id="w-node-d99af7bd5692-e21305a8">Dentissimo содержит натуральные ингредиенты: Экстракт Герани Макулатум, Эхинацея узколистная, Ромашка аптечная, Шалфей лекарственный, Цетрария исландская, Папаин, Мята перечная, экстракт Алоэ, Эвкалипт, масло Чайного дерева, экстракт Крамерии, масло Комиффоры мирровой, Тимол.<br/><br/>Вы найдете полный список ингредиентов на обратной стороне упаковки и тубы, а также в описании товара.</p></div></div></div></section><div class="separator"><img src="https://uploads-ssl.webflow.com/5e4a6a8dd9eb2c003ad03de7/5e4a984373dcf3e2f6197e55_full_logo_white.png" alt="" class="image"/></div><div id="contact" class="contact"><div class="container-2 w-container"><h2>ООО «Денталик»</h2><p class="paragraph-2">97198 г. Санкт-Петербург, ул. Ропшинская д 4 лит. А пом. 2Н<br/>График работы: пн-пт с 10.00 до 18.00<br/></p><div class="div-block-2"><div class="div-block"><img src="https://uploads-ssl.webflow.com/5e4e4fa75e5be328c3a7fad2/5e4e50f62616d3275da48c62_phone.png" height="40" width="40" alt="" class="image-2"/><a href="tel:+78123808290" class="link-2">+7 (812) 380-82-90</a></div><div class="div-block"><img src="https://uploads-ssl.webflow.com/5e4e4fa75e5be328c3a7fad2/5e4e50f60f3ef814c161b216_mail.png" height="40" width="40" alt="" class="image-2"/><a href="mailto:skindoc@mail.ru" class="link-3">skindoc@mail.ru</a></div></div><div class="yandex_map"></div><div class="text-block">Copyright © 2020 Dentalik. All rights reserved.</div></div></div><div class="btn_up"><a href="#menu" class="link">↑</a></div><script src="https://d3e54v103j8qbb.cloudfront.net/js/jquery-3.4.1.min.220afd743d.js" type="text/javascript" integrity="sha256-CSXorXvZcTkaix6Yvo6HppcZGetbYMGWSFlBw8HfCJo=" crossorigin="anonymous"></script><script src="https://uploads-ssl.webflow.com/5e4e4fa75e5be328c3a7fad2/js/webflow.d04fea9ea.js" type="text/javascript"></script><!--[if lte IE 9]><script src="//cdnjs.cloudflare.com/ajax/libs/placeholders/3.0.2/placeholders.min.js"></script><![endif]--><!-- FOOTER CODE --></body></html>`

	var re = regexp.MustCompile(`(?m)<script src="https://uploads-ssl\.webflow\.com.*js" type="text/javascript"></script>`)
	for i, match := range re.FindAllString(str, -1) {
		fmt.Println(match, "found at index", i)
	}

	re = regexp.MustCompile(`(?m)<link href=.* rel="stylesheet" type="text\/css"\/>`)
	for i, match := range re.FindAllString(str, -1) {
		fmt.Println(match, "found at index", i)
	}

}
