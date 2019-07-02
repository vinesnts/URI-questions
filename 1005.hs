import System.IO
import Text.Printf

media :: Float -> Float -> Float 
media a b = (((a * 3.5) + (b * 7.5)) / (3.5 + 7.5))

main = do 
       a <- readLn 
       b <- readLn 
       printf "MEDIA = %0.5f\n" (media a b)
