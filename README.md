# url shorty
- a challenge by mi amigo

## requirements:
- [x] written in golang
- [x] return short url for a given long url & return long url for a given short url  (bijective function)
- [x] ability for the user to delete a short url

- [ ] HONORARY: use a database instead of local mappings

## approach:
```
take a long url -> hash it -> store in the l2s map -> while simultaneously storing the reverse mapping in the s2l map 
```
