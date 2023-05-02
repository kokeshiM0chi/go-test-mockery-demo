

hello:
	mockery --dir=internal/repositories \
	--name=ProductRepositoryInterface \
	--filename=product_repository_interface.go \
	--output=internal/mocks/repomocks \
	--outpkg=repomocks \

# リポジトリ直下にできる
hello2:
	mockery --all