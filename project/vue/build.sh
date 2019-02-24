#!/bin/bash

# /css /js を/static/products/vocabulary以下に変更
sed -i s?/css?/static/products/vocabulary/css?g dist/index.html
sed -i s?/js?/static/products/vocabulary/js?g dist/index.html
sed -i s?/img?/static/products/vocabulary/img?g dist/index.html
