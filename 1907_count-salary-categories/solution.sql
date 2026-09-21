with stats as (
    select
        sum(if(a.income < 20000, 1, 0)) as low,
        sum(if(20000 <= a.income and a.income <= 50000, 1, 0)) as mid,
        sum(if(a.income > 50000, 1, 0)) as high
    from Accounts a
)
select
    'Low Salary' category,
    stats.low accounts_count
from stats
union (
    select
        'Average Salary',
        stats.mid accounts_count
    from stats
)
union (
    select
        'High Salary',
        stats.high accounts_count
    from stats
)
