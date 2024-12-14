DELETE FROM tenant;
DELETE FROM prescription;

INSERT INTO
  tenant (id, key)
VALUES
  (1, "foobar");

INSERT INTO
  prescription (id, tenant, interval, interval_unit, start_date, medicine, amount, unit)
VALUES
  (1, 1, 1, "daily", "2024-11-29", "Panadol", 500, "mg");

INSERT INTO
  serving (id, tenant, prescription, medicine, amount, unit, taken, scheduled_at, taken_at)
VALUES
  (1, 1, 1, "Panadol", 500, "mg", TRUE, "2024-11-29T08:00:00", "2024-11-29T08:04:12");
