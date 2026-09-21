select
    p.product_id,
    coalesce(max(sub.new_price), 10) as price
from Products p
left outer join (
    select
        p.product_id,
        p.new_price,
        p.change_date,
        row_number() over (partition by p.product_id order by p.change_date desc) rn
    from Products p
    where p.change_date <= '2019-08-16'
) sub on p.product_id = sub.product_id and sub.rn = 1
group by p.product_id
;