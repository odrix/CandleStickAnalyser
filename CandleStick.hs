{-# LANGUAGE OverloadedStrings #-}
{-# LANGUAGE DeriveGeneric #-}
{-# OPTIONS_GHC -Wno-incomplete-patterns #-}

module CandleStick where

import GHC.Generics

data CandleStick = CandleStick {
    openTime :: Integer,
    open    :: Double,
    hight   :: Double,
    low     :: Double,
    close   :: Double,
    volume  :: Double,
    closeTime :: Integer
} deriving (Generic, Eq, Show)

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

lowerShadowRatio :: CandleStick -> Double
lowerShadowRatio candle  = lowerShadow candle / body candle

upperShadowRatio :: CandleStick -> Double
upperShadowRatio candle = upperShadow candle / body candle

isRed :: CandleStick -> Bool
isRed candle = open candle > close candle

isGreen :: CandleStick -> Bool
isGreen candle = open candle < close candle

bodyPourcent :: CandleStick -> Double
bodyPourcent candle = abs (open candle - close candle) / open candle

hasSmallBody :: CandleStick -> Bool
hasSmallBody candle = bodyPourcent candle < maxSmallBody

hasNormalBody :: CandleStick -> Bool
hasNormalBody candle = bodyPourcent candle > maxSmallBody

hasBigBody :: CandleStick -> Bool
hasBigBody candle =  bodyPourcent candle > minBigBody

hasCloseBody :: CandleStick -> Bool
hasCloseBody candle = bodyPourcent candle < maxCloseBody


isHammer :: CandleStick -> Bool
isHammer candle
    | isGreen candle && hasSmallBody candle && lowerShadowRatio candle > minLongShadow && upperShadowRatio candle < maxTailShadow = True
    | isRed candle && hasSmallBody candle && upperShadowRatio candle > minLongShadow && lowerShadowRatio candle < maxTailShadow = True
    | otherwise = False

isDoji :: CandleStick -> Bool
isDoji = hasCloseBody


isMarubozu :: CandleStick -> Bool
isMarubozu candle = hasNormalBody candle && upperShadowRatio candle < maxTailShadow && lowerShadowRatio candle < maxTailShadow

isMorningStar :: [CandleStick] -> Bool
isMorningStar [a,b,c] = isRed a && isDoji b && isGreen c


maxCloseBody = 0.005
maxSmallBody = 0.02
minBigBody = 0.05
maxTailShadow = 0.05
minLongShadow = 0.6