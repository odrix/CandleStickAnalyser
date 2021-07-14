
module BinanceAdapter where

import CandleStick

readCandleStick :: [String] -> CandleStick
readCandleStick [] =  CandleStick 0 0 0 0 0 0 0
readCandleStick [a] = CandleStick (read a :: Integer) 0 0 0 0 0 0
readCandleStick [a, b] = CandleStick (read a :: Integer) (read b :: Double) 0 0 0 0 0
readCandleStick [a, b, c] = CandleStick (read a :: Integer) (read b :: Double) (read c :: Double) 0 0 0 0
readCandleStick [a, b, c, d] = CandleStick (read a :: Integer) (read b :: Double) (read c :: Double) (read d :: Double) 0 0 0
readCandleStick [a, b, c, d, e] = CandleStick (read a :: Integer) (read b :: Double) (read c :: Double) (read d :: Double) (read e :: Double) 0 0
readCandleStick [a, b, c, d, e, f] = CandleStick (read a :: Integer) (read b :: Double) (read c :: Double) (read d :: Double) (read e :: Double)  (read f :: Double) 0
readCandleStick (a:b:c:d:e:f:g:xs) =
    let openTime = read a :: Integer
        open = read b :: Double
        hight = read c :: Double
        low = read d :: Double
        close =  read e :: Double
        volume =  read f :: Double
        closeTime = read g :: Integer
    in CandleStick openTime open hight low close volume closeTime


readCandleStickList :: [[String]] -> [CandleStick]
readCandleStickList = map readCandleStick
