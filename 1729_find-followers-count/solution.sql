select
    f.user_id,
    count(f.follower_id) followers_count
from Followers f
group by f.user_id
order by f.user_id
;
