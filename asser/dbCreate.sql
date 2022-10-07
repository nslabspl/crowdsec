DECLARE
	-- IP address
	ip float;
	-- Next IP in the range
	next_ip float;
	-- Previous IP
	prev_ip float;
	-- Numbver of IPs in sequence
	ip_seq float;
	-- ID of specific IP
	ip_id integer;
begin
  function checkCorrectness(
	ip in float
  ) return dbms_output.put_line('IP is of correct format.');

  procedure addToCTI(
	for ip in ip_seq
	insert into IP.curated (ipnumber)
	values (ip)
	end
  );
end;