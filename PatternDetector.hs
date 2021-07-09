module PatternDetector where

import CandleStick

detect :: [CandleStick] -> [String]
detect [] = []
detect (x:xs)
    | isHammer x = "hammer" : detect xs
    | isDoji x = "doji": detect xs
    | isMarubozu x = "marubozu" : detect xs
    | otherwise = "" : detect xs

mtl :: Maybe [a] -> [a]
mtl Nothing = []
mtl (Just [xs]) = [xs]