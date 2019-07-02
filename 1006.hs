import System.IO
import Text.Printf

media :: Float -> Float -> Float -> Float 
media a b c = (((a * 2) + (b * 3) + (c * 5)) / (2 + 3 + 5))

main = do 
       a <- readLn 
       b <- readLn 
       c <- readLn
       printf "MEDIA = %0.1f\n" (media a b c)
