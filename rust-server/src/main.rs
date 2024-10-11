#[macro_use]
extern crate rocket;

#[get("/")]
fn index() -> &'static str {
    "Hello, world!"
}
use rocket::serde::json::Json;
use redis::Commands;
use serde::{Serialize, Deserialize};
use bit_set::BitSet;
use serde_json::value;

#[derive(Serialize, Deserialize)]
struct RedisResponse {
    message: String,
}

// example redis url: redis://<host>:<port>
static REDIS_URL: &str = "redis://:pando.dev@localhost:6379";

#[get("/redis/<key>")]
async fn get_redis_value(key: String) -> Json<RedisResponse> {
    let client = redis::Client::open(REDIS_URL).unwrap();
    let mut con = client.get_connection().unwrap();
    
    let mut bits: BitSet = BitSet::with_capacity(1000000);
    let bv = bits.get_mut();
    print!("---10: {}", bv[10]);
    print!("---100: {}", bv[100]);
    
    bv.set(10, true);
    print!("---10: {}", bv[10]);

    // con.set("bitset", bv.get_mut(index)).unwrap();

    match con.get::<_, String>(key) {
        Ok(value) => Json(RedisResponse { message: format!("ok") }),
        // Err(_) => Json(RedisResponse { message: "Key not found".to_string() }),
        Err(_) => Json(RedisResponse { message: format!("{}", bv.capacity()) }),
    }
}

#[get("/redis/<key>/<value>")]
async fn set_redis_value(key: String, value: String) -> Json<RedisResponse> {
    let client = redis::Client::open(REDIS_URL).unwrap();
    let mut con = client.get_connection().unwrap();
    match con.set::<_, String, String>(key, value) {
        Ok(_) => Json(RedisResponse { message: "Value set".to_string() }),
        Err(_) => Json(RedisResponse { message: "Failed to set value".to_string() }),
    }
}

#[launch]
fn rocket() -> _ {
    rocket::build()
        .mount("/", routes![index, get_redis_value, set_redis_value])
}