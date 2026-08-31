select users.name, sum(coalesce(rides.distance, 0)) travelled_distance 
from users
left join rides on users.id = rides.user_id
group by users.name, users.id
order by travelled_distance desc, users.name asc;