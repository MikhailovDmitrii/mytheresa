build:
	docker build -t mikhailov-promotions .

run:
	docker run -u mytheresauser -it -v ${PWD}/var/products.db:/app/var/products.db -p 8080:8080 mikhailov-promotions -d

test:
	docker build -f Dockerfile-test --progress plain --no-cache .


# It sends a request to insert a sample of 5 products to database
insert:
	curl -v -X POST "http://127.0.0.1:8080/products" -d "@./products.json"