select
    date_format(t.trans_date, '%Y-%m') as month,
    t.country,
    count(t.id) as trans_count,
    sum(if(t.state = 'approved', 1, 0)) as approved_count,
    sum(t.amount) as trans_total_amount,
    sum(if(t.state = 'approved', t.amount, 0)) as approved_total_amount
from Transactions t
group by month, country
;
