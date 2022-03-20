
app_image := "my-go-webapp"
port_mapping := "8080:8081"

build:
	@echo "build......"
	docker build -t ${app_image} .

run: build
	docker run -it -p ${port_mapping} ${app_image}

run_detached: build
	docker run -d -p ${port_mapping} ${app_image}

stop:
	docker stop $(docker ps | grep ${app_image} | cut -d " " -f1)

clean:
	rm -rf 	main


