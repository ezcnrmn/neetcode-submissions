select first_name, last_name, city, state 
from person p
left join address a using(person_id)
