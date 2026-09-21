select
    t.x,
    t.y,
    t.z,
    if(t.x+t.y>t.z and t.y+t.z>t.x and t.z+t.x>t.y, 'Yes', 'No') triangle
from Triangle t;
