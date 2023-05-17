-- Генерируем университеты и направления
WITH names AS (
    SELECT unnest('{"
Massachusetts Institute of Technology (MIT) ","University of Cambridge ","
Stanford University ","University of Oxford ","Harvard University ","California Institute of Technology (Caltech) ","Imperial College London ","UCL (University College London) ",
                "
ETH Zurich (Swiss Federal Institute of Technology)"}'::text[]) names
)

   ,facilities AS (
    SELECT unnest('{"math","philosophy","it","computer science","data science", "physics","PE","management"}'
        ::text[]) facilities
)
insert into university (univ_name,facility)
SELECT  *
FROM    names
            FULL JOIN  facilities on true;

-- Генерируем имена студентов, их оценки и университеты, в которые они подали бумаги
WITH names AS (
    SELECT unnest('{"Smith","Johnson","Williams","Brown","Jones","Garcia","Miller","Davis",
                "Rodriguez","Martinez","Hernandez","Lopez","Gonzalez",
                "Wilson","Anderson","Thomas","Taylor","Moore","Jackson","Martin","Lee",
                "Perez","Thompson","White","Harris","Sanchez","Clark",
                "Ramirez","Lewis","Robinson","Walker","Young","Allen","King",
                "Wright","Scott","Torres","Nguyen","Hill","Flores","Green","Adams",
                "Nelson","Baker","Hall","Rivera","Campbell","Mitchell","Carter","Roberts"}'::text[]) name
)
   ,grades AS (
    SELECT unnest(ARRAY[generate_series(120, 315)]) grade
)
   , univ_ids as (
    select unnest(array[generate_series(1, 75)])
)
insert into student (name, grades, univ_apply_id)
SELECT  *
FROM    names
            FULL JOIN  grades on true join univ_ids on true;


delete from student where true;
delete from university where true;

