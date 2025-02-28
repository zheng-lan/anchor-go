build:
	go build && rm -rf ./generated

swap-proxy:build
	./anchor-go -type-id=uint8 -src=./swap_proxy.json -pkg=swap_proxy -dst=./generated/swap_proxy

temp:build
	./anchor-go -type-id=uint8 -src=./temp.json -pkg=temp -dst=./generated/temp