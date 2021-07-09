{-# LANGUAGE OverloadedStrings #-}
{-# LANGUAGE DeriveGeneric #-}

import Data.Aeson
import qualified Data.ByteString.Lazy as B
import GHC.Generics
import PatternDetector
import CandleStick

instance FromJSON CandleStick

main = do
  input <- B.readFile "klines_BTCUSDT_H1.json"
  let mm = decode input :: Maybe [CandleStick]
  case mm of
    Nothing -> print "error parsing JSON"
    Just m -> (putStrLn.greet) m
    
greet m = (show.name) m ++" was born in the year "++ (show.born) m