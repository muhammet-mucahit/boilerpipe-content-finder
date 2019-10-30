# python-boilerpipe

## Configuration

Dependencies:

 * jpype
 * chardet
 * fastapi[all]
 * hypercorn
 * beautifulsoup4

The boilerpipe jar files will get fetched and included automatically when building the package.

## Installation

Checkout the code:

	git clone https://github.com/muhammet-mucahit/boilerpipe-content-finder
	cd boilerpipe-content-finder/python-boilerpipe

**virtualenv**
```code
	virtualenv env
	source env/bin/activate
    pip install -r requirements.txt
	python setup.py install
```

## Running

	hypercorn main:app --bind localhost:5000 --reload

## Request

	curl -X POST \
		http://127.0.0.1:5000/api/v1/boilerpipe \
		-H 'Accept: */*' \
		-H 'Accept-Encoding: gzip, deflate' \
		-H 'Cache-Control: no-cache' \
		-H 'Connection: keep-alive' \
		-H 'Content-Length: 1025' \
		-H 'Content-Type: application/json' \
		-H 'Host: 127.0.0.1:5000' \
		-H 'Postman-Token: 1c460376-e3d6-429f-948a-81bb5e10f781,eb364200-c29c-461f-af15-84a60d55c03b' \
		-H 'User-Agent: PostmanRuntime/7.18.0' \
		-H 'cache-control: no-cache' \
		-d '{
			"urls": [
				"https://zeo.org/tr/seo/",
				"https://www.mediamarkt.com.tr/tr/product/_%C4%B1deapad-s145-amd-ryzen5-3500u-3-7ghz-8gb-ram-512gb-ssd-vega8-15-6-windows-10-81ut0017tx-laptop-1203872.html",
				"https://www.hepsiburada.com/samsung-galaxy-a20-32-gb-samsung-turkiye-garantili-p-HBV00000JFE8V?magaza=CepteTek&wt_gl=cpc.6804.shop.elk.telefon-cep-telefonu-ssc&gclid=EAIaIQobChMIxJ7P5_O05QIVjRnTCh39ZwPNEAQYASABEgLSVfD_BwE",
				"https://github.com/misja/python-boilerpipe",
				"https://simpleisbetterthancomplex.com/tutorial/2018/11/22/how-to-implement-token-authentication-using-django-rest-framework.html",
				"https://katemats.com/what-every-programmer-should-know-about-seo/",
				"https://chriskiehl.com/article/parallelism-in-one-line",
				"https://medium.com/backticks-tildes/lets-build-an-api-with-django-rest-framework-32fcf40231e5",
				"https://golang.org/pkg/strings/",
				"https://medium.com/t%C3%BCrkiye/seo-bu-kadar-%C3%B6nemli-mi-56f11fe3754d"
			]
		}'

