#!/bin/bash

rm -rf github.com/bmizerany
rm -rf github.com/google
rm -rf code.google.com

git clone https://github.com/bmizerany/pat.git github.com/bmizerany/pat
rm -rf github.com/bmizerany/pat/.git
rm -rf github.com/bmizerany/pat/example

git clone git@github.com:google/go-querystring.git github.com/google/go-querystring
rm -rf github.com/google/go-querystring/.git

git clone git@github.com:google/go-github.git github.com/google/go-github
rm -rf github.com/google/go-github/.git

hg clone https://code.google.com/p/goauth2/ code.google.com/p/goauth2
rm -rf code.google.com/p/goauth2/appengine
rm -rf code.google.com/p/goauth2/lib
rm -rf code.google.com/p/goauth2/oauth/jwt
rm -rf code.google.com/p/goauth2/oauth/example
