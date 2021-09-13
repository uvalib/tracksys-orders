#!/bin/bash
go run backend/*.go -port 8085  \
   -dbhost localhost \
   -dbname tracksys_prod \
   -dbuser root \
   -dbpass pass 


#   -dbhost db1.lib.virginia.edu \
#   -dbname tracksys_production \
#   -dbuser tracksys_ro \
#   -dbpass F75jrup8
