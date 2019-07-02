import System.IO

soma :: Integer -> Integer -> Integer 
soma a b = a + b

main = do 
       a <- readLn 
       b <- readLn 
       putStr $ "SOMA = "; print(soma a b)
