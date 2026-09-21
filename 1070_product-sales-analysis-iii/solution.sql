select
    s.product_id,
    s.year first_year,
    s.quantity,
    s.price
from (
    select
        product_id,
        year,
        quantity,
        price,
        rank() over (partition by product_id order by year) rk
    from Sales
) as s
where s.rk = 1
;
