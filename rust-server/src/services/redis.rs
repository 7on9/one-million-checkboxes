use redis::{Commands, Client, RedisResult};

pub struct RedisService {
  client: Client,
}

impl RedisService {
  pub fn new(url: &str) -> RedisResult<Self> {
    let client = Client::open(url)?;
    Ok(RedisService { client })
  }

  pub fn set_value(&self, key: &str, value: &str) -> RedisResult<()> {
    let mut con = self.client.get_connection()?;
    con.set(key, value)?;
    Ok(())
  }

  pub fn get_value(&self, key: &str) -> RedisResult<String> {
    let mut con = self.client.get_connection()?;
    let value: String = con.get(key)?;
    Ok(value)
  }
}

fn connect() -> redis::Connection {
  //format - host:port
  let redis_host_name =
      env::var("REDIS_HOSTNAME").expect("missing environment variable REDIS_HOSTNAME");
  
  let redis_password = env::var("REDIS_PASSWORD").unwrap_or_default();
  //if Redis server needs secure connection
  let uri_scheme = match env::var("IS_TLS") {
      Ok(_) => "rediss",
      Err(_) => "redis",
  };
  let redis_conn_url = format!("{}://:{}@{}", uri_scheme, redis_password, redis_host_name);
  redis::Client::open(redis_conn_url)
      .expect("Invalid connection URL")
      .get_connection()
      .expect("failed to connect to Redis")
}
