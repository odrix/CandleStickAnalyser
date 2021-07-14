module TestBinanceAdapter where

import Test.HUnit
import BinanceAdapter
import CandleStick

testReadCandleStickForEmptyList :: Test
testReadCandleStickForEmptyList = 
    TestCase $ assertEqual "Should return default candle for empty list"
                           (CandleStick 0 0 0 0 0 0 0) (readCandleStick ([]::[String]))

testReadCandleStickOneValueList :: Test
testReadCandleStickOneValueList =
    TestCase $ assertEqual "Should return candle with openTime"
                            (CandleStick 1 0 0 0 0 0 0) (readCandleStick ["1"])

testReadCandleStickSevenValueList :: Test
testReadCandleStickSevenValueList =
    TestCase $ assertEqual "Should return candle filled"
                            (CandleStick 1 2 3 4 5 6 7) (readCandleStick ["1", "2", "3", "4", "5", "6", "7"])


testReadCandleStickListTwoValues :: Test
testReadCandleStickListTwoValues =
    TestCase $ assertEqual "Should return 2 candles filled"
                            ([(CandleStick 1 2 3 4 5 6 7), (CandleStick 9 8 7 6 5 4 3)]) (readCandleStickList [["1", "2", "3", "4", "5", "6", "7"], ["9", "8", "7", "6", "5", "4", "3"]])

main :: IO Counts
main = runTestTT $ TestList [testReadCandleStickForEmptyList, testReadCandleStickOneValueList, testReadCandleStickSevenValueList, testReadCandleStickListTwoValues]

