export MG_API_KEY = 
export MG_DOMAIN =
export MG_PUBLIC_API_KEY = pubkey-86b3b4f73639db16b3bbd7c6a19b43ad
export MG_URL = 
export TEMPLATES_FOLDER = templates



build:
	go build

dev:
	./site

prod:
	nohup ./site

kill:
	pkill site
