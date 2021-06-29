data CandleStick = CandleStick {
    open    :: Double,
    close   :: Double,
    hight   :: Double,
    low     :: Double
} deriving (Show)

body :: CandleStick -> Double
body candle = abs (open candle - close candle)

lowerShadow :: CandleStick -> Double
lowerShadow candle 
    | isRed candle = low candle - close candle
    | otherwise = low candle - open candle

upperShadow :: CandleStick -> Double
upperShadow candle 
    | isRed candle = hight candle - open candle
    | otherwise = hight candle - close candle

isRed :: CandleStick -> Bool
isRed candle = open candle > close candle

isGreen :: CandleStick -> Bool
isGreen candle = open candle < close candle
