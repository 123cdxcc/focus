# 查询总量
```sql
select count(id) as total from p_user;
```
# 查询用户分页及最新创建订单信息
```sql
select id,created_at,updated_at,nickname,mobile,
(select p_order.product_name from p_order where p_user.id=p_order.user_id order by created_at desc limit 1) as product_name,
(select p_order.total_price from p_order where p_user.id=p_order.user_id order by created_at desc limit 1) as total_price,
(select p_order.count from p_order where p_user.id=p_order.user_id order by created_at desc limit 1) as count,
(select p_order.unit_price from p_order where p_user.id=p_order.user_id order by created_at desc limit 1) as unit_price,
(select p_order.pay_type from p_order where p_user.id=p_order.user_id order by created_at desc limit 1) as pay_type,
(select p_order.status from p_order where p_user.id=p_order.user_id order by created_at desc limit 1) as status,
(select p_order.created_at from p_order where p_user.id=p_order.user_id order by p_order.created_at desc limit 1) as create_at
from p_user
limit ?, ?;
```