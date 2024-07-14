CREATE EVENT daily_snapshot_event
ON SCHEDULE
    EVERY 1 DAY
    STARTS CURRENT_DATE() + INTERVAL 1 DAY
DO
    insert into qrmark_snapshots (school_id, company_id, total_points, snapshot_date)
    select
        q.school_id,
        q.company_id,
        sum(m.points) as total_points,
        now() as snapshot_date
    from
        (
            select distinct
                school_id,
                company_id
            from
                qrmarks
            where
                created_at >= coalesce((select max(snapshot_date) from qrmark_snapshots where school_id = 7026), '1970-01-01')
        ) as q
        join qrmarks as m on q.school_id = m.school_id and q.company_id = m.company_id
    group by
        q.school_id, q.company_id;
