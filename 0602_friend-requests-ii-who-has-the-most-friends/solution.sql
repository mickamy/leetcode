with all_relationships as (
    select requester_id as user_id from RequestAccepted
    union all
    select accepter_id as user_id from RequestAccepted
)
select
    user_id as id,
    count(*) as num
from all_relationships
group by 1
order by 2 desc
limit 1
;
