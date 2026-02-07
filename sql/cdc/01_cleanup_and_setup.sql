-- 1. Drop old/zombie slots to free up max_wal_senders
-- Intentionally using separate calls as specific slots may or may not exist
SELECT pg_drop_replication_slot('cdc_slot_bets') WHERE EXISTS (SELECT 1 FROM pg_replication_slots WHERE slot_name = 'cdc_slot_bets');
SELECT pg_drop_replication_slot('cdc_slot_race') WHERE EXISTS (SELECT 1 FROM pg_replication_slots WHERE slot_name = 'cdc_slot_race');
SELECT pg_drop_replication_slot('cdc_slot_board') WHERE EXISTS (SELECT 1 FROM pg_replication_slots WHERE slot_name = 'cdc_slot_board');
SELECT pg_drop_replication_slot('cdc_slot_process') WHERE EXISTS (SELECT 1 FROM pg_replication_slots WHERE slot_name = 'cdc_slot_process');
SELECT pg_drop_replication_slot('cdc_slot_retired') WHERE EXISTS (SELECT 1 FROM pg_replication_slots WHERE slot_name = 'cdc_slot_retired');
SELECT pg_drop_replication_slot('cdc_slot_activations') WHERE EXISTS (SELECT 1 FROM pg_replication_slots WHERE slot_name = 'cdc_slot_activations');
SELECT pg_drop_replication_slot('cdc_slot') WHERE EXISTS (SELECT 1 FROM pg_replication_slots WHERE slot_name = 'cdc_slot');
SELECT pg_drop_replication_slot('cdc_slot_main') WHERE EXISTS (SELECT 1 FROM pg_replication_slots WHERE slot_name = 'cdc_slot_main');

-- 2. Drop old publication if exists to start fresh
DROP PUBLICATION IF EXISTS corridor_cdc_pub;
DROP PUBLICATION IF EXISTS betting_cdc_pub;

-- 3. Create unified publication
CREATE PUBLICATION betting_cdc_pub FOR TABLE 
    gaming.bet, 
    gaming.races_process_groups, 
    gaming.board_race_group, 
    gaming.retired_horse_group, 
    gaming.race, 
    gaming.group_race_activations;

-- 4. Create the single replication slot
-- We use logical_decoding with pgoutput plugin
SELECT pg_create_logical_replication_slot('betting_main_slot', 'pgoutput');
