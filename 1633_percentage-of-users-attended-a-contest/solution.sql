select
    r.contest_id,
    round(count(r.user_id)/max(counts.user)*100, 2) as percentage
from Register r
left outer join Users u on r.user_id = u.user_id
cross join (
    select count(distinct user_id) user
    from Users
) as counts
group by r.contest_id
order by percentage desc, r.contest_id
;
