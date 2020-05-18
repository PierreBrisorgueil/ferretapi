# Docker [Ferret](https://github.com/MontFerret/ferret) API
<p align="center">
	<a href="https://goreportcard.com/report/github.com/PierreBrisorgueil/ferretApi">
		<img alt="Go Report Status" src="https://goreportcard.com/badge/github.com/PierreBrisorgueil/ferretApi">
	</a>
	<a href="https://discord.gg/kzet32U">
		<img alt="Discord Chat" src="https://img.shields.io/discord/501533080880676864.svg">
	</a>
	<a href="https://github.com/PierreBrisorgueil/ferretApi/releases">
		<img alt="Ferret Api release" src="https://img.shields.io/github/release/PierreBrisorgueil/ferretApi.svg">
	</a>
	<a href="http://opensource.org/licenses/MIT">
		<img alt="MIT License" src="http://img.shields.io/badge/license-MIT-brightgreen.svg">
	</a>
</p>

![ferret](https://raw.githubusercontent.com/MontFerret/ferret/master/assets/intro.jpg)

## What is [Ferret](https://github.com/MontFerret/ferret) ?
[```ferret```](https://github.com/MontFerret/ferret) is a web scraping system. It aims to simplify data extraction from the web for UI testing, machine learning, analytics and more.    
[```ferret```](https://github.com/MontFerret/ferret) allows users to focus on the data. It abstracts away the technical details and complexity of underlying technologies using its own declarative language. 
It is extremely portable, extensible and fast.

[Read the introductory blog post about Ferret here!](https://medium.com/@ziflex/say-hello-to-ferret-a-modern-web-scraping-tool-5c9cc85ba183)

## What is this container  ?

This container makes it possible to deploy an instance including an API and chrome headless. You can query this API by sending an FQL instructions via a POST request. This request will be executed via Ferret & Chrome and the result will be returned to you.

## Show me some code

```
curl -d "{\"query\": \"LET doc = DOCUMENT('https://weareopensource.me', true) LET btn = ELEMENT(doc, '.nav-hobbies') CLICK(btn) WAIT_NAVIGATION(doc) FOR el IN ELEMENTS(doc, '.post-card-title') RETURN TRIM(el.innerText)\"}" -H "Content-Type: application/json" -X POST http://localhost:8080/
```

## Installation

### Docker Hub

```
docker run --rm -p 8080:8080 pierrebrisorgueil/ferretapi
```

### Build
```
git clone && cd ferretApi
docker build -t ferretapi .
docker run --rm -p 8080:8080 ferretapi
```

## Dev

Pierre 

[![Github](https://badges.weareopensource.me/badge/Follow-me%20on%20Github-282828.svg?style=flat-square)](https://github.com/PierreBrisorgueil) [![Twitter](https://badges.weareopensource.me/badge/Follow-me%20on%20Twitter-3498db.svg?style=flat-square)](https://twitter.com/pbrisorgueil?lang=fr) [![Youtube](https://badges.weareopensource.me/badge/Watch-me%20on%20Youtube-e74c3c.svg?style=flat-square)](https://www.youtube.com/channel/UCIIjHtrZL5-rFFupn7c3OtA) [![Instagram](https://badges.weareopensource.me/badge/Follow-me%20on%20Instagram-f27231.svg?style=flat-square)](https://www.instagram.com/pierre_brsrgl/) [![Linkedin](https://badges.weareopensource.me/badge/Add-me%20on%20linkedin-006DA9.svg?style=flat-square)](https://www.linkedin.com/in/pierre-brisorgueil/)

