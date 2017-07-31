-- Function: public.yqueue_del(integer)

-- DROP FUNCTION public.yqueue_del(integer);

CREATE OR REPLACE FUNCTION public.yqueue_del(p_id integer)
  RETURNS void AS
$BODY$
	delete from yqueue where id = p_id
$BODY$
  LANGUAGE sql VOLATILE
  COST 100;
ALTER FUNCTION public.yqueue_del(integer)
  OWNER TO postgres;


-----------------

-- Function: public.yqueue_get(integer)

-- DROP FUNCTION public.yqueue_get(integer);

CREATE OR REPLACE FUNCTION public.yqueue_get(p_id integer)
  RETURNS SETOF yqueue AS
$BODY$
	select * from yqueue where id = (case when p_id is null then id else p_id end)
	order by lastupdate asc
$BODY$
  LANGUAGE sql VOLATILE
  COST 100
  ROWS 1000;
ALTER FUNCTION public.yqueue_get(integer)
  OWNER TO postgres;

----------

-- Function: public.yqueue_get(integer, integer)

-- DROP FUNCTION public.yqueue_get(integer, integer);

CREATE OR REPLACE FUNCTION public.yqueue_get(
    p_id integer,
    p_status integer, 
    p_who text)
  RETURNS SETOF yqueue AS
$BODY$
	select * from yqueue 
	where 
		id = (case when p_id is null then id else p_id end) 
		and 
		status = (case when p_status is null then status else p_status end)
    and 
    who = (case when p_who is null then who else p_who end)
	order by lastupdate asc
$BODY$
  LANGUAGE sql VOLATILE
  COST 100
  ROWS 1000;
ALTER FUNCTION public.yqueue_get(integer, integer, text)
  OWNER TO postgres;

----------

-- Function: public.yqueue_ins(text)

-- DROP FUNCTION public.yqueue_ins(text);

CREATE OR REPLACE FUNCTION public.yqueue_ins(
  p_yturl text,
  p_ytkey text,
  p_who text)
  RETURNS integer
LANGUAGE 'sql'
AS $BODY$
  insert into yqueue (yturl, ytkey, status, file, who, lastupdate) values (p_yturl, p_ytkey, 1, null, p_who, now())
  returning id
$BODY$;

ALTER FUNCTION public.yqueue_ins(text, text)
  OWNER TO postgres;

----------

-- Function: public.yqueue_upd(integer, integer, text)

-- DROP FUNCTION public.yqueue_upd(integer, integer, text);

CREATE OR REPLACE FUNCTION public.yqueue_upd(
    p_id integer,
    p_status integer DEFAULT NULL::integer,
    p_file text DEFAULT NULL::text)
  RETURNS void AS
$BODY$
	update yqueue set 
		status = case when p_status is null then status else p_status end, 
		file = case when p_file is null then file else p_file end, 
		lastupdate = now()
	where id = p_id
$BODY$
  LANGUAGE sql VOLATILE
  COST 100;
ALTER FUNCTION public.yqueue_upd(integer, integer, text)
  OWNER TO postgres;


----------

-- Function: public.yqueue_get_users()

-- DROP FUNCTION public.yqueue_get_users();

CREATE OR REPLACE FUNCTION public.yqueue_get_users()
  RETURNS TABLE(who text) AS
$BODY$
	select distinct who from yqueue where status = 3
$BODY$
  LANGUAGE sql VOLATILE
  COST 100
  ROWS 1000;
ALTER FUNCTION public.yqueue_get_users()
  OWNER TO postgres;
