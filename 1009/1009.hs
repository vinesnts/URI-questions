import System.IO
import Text.Printf

totalSalario :: String -> Float -> Float -> Float 
totalSalario nome a b = a + b * 0.15

main = do 
       nome <- getLine 
       a <- readLn 
       b <- readLn
       printf "TOTAL = %0.2f\n" (totalSalario nome a b)
