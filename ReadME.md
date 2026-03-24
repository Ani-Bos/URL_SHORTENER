overall design is take actual url do a base62 hash of it generate a shorturl and then get it
But there can be collision as well whether that id is present or not if present 
use a bloom filter to check if that id is already present if present add a predefine value to thw actual_url generATE A HASH
//or redis can also work if that id is present as key sp add a predefines tring

Bas62 conversion

it depends on static id generator so collision is not possible at all
security vulnerability as we can predict the next id not so scalable

hash+collision resolution

fixed url short length no requirement of static ids
collison is there so need to be resolved
also not possible to predict next url as it doesn't depend on ids so random code generation
much of scalable idea