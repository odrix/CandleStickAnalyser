import CandleStick

detect :: [CandleStick] -> [String]
detect [] = []
detect (x:xs)
    | isHammer x = "hammer" : detect xs
    | isDoji x = "doji": detect xs
    | isMarubozu x = "marubozu" : detect xs
    | otherwise = "" : detect xs