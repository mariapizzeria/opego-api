--
-- PostgreSQL database dump
--

-- Dumped from database version 17.5
-- Dumped by pg_dump version 17.5

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: opego; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE opego WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'ru';


ALTER DATABASE opego OWNER TO postgres;

\connect opego

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: driver; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.driver (
    driver_id integer NOT NULL,
    name character varying(255) NOT NULL,
    car_type text NOT NULL,
    car_number character varying(255),
    score numeric(3,2),
    available boolean,
    current_location jsonb
);


ALTER TABLE public.driver OWNER TO postgres;

--
-- Name: driver_driver_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.driver_driver_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.driver_driver_id_seq OWNER TO postgres;

--
-- Name: driver_driver_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.driver_driver_id_seq OWNED BY public.driver.driver_id;


--
-- Name: order; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."order" (
    order_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    canceled_at timestamp without time zone,
    passenger_id integer,
    order_status character varying(255),
    driver_assigned integer,
    arrived_code character varying(255) DEFAULT NULL::character varying,
    address_from text,
    address_to text,
    tariff character varying(255),
    selected_services text[],
    comment text,
    price numeric,
    completed_at timestamp without time zone
);


ALTER TABLE public."order" OWNER TO postgres;

--
-- Name: order_order_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.order_order_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.order_order_id_seq OWNER TO postgres;

--
-- Name: order_order_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.order_order_id_seq OWNED BY public."order".order_id;


--
-- Name: passenger; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.passenger (
    passenger_id integer NOT NULL,
    name character varying(255),
    city character varying(255),
    phone_number character varying(255)
);


ALTER TABLE public.passenger OWNER TO postgres;

--
-- Name: passenger_passenger_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.passenger_passenger_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.passenger_passenger_id_seq OWNER TO postgres;

--
-- Name: passenger_passenger_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.passenger_passenger_id_seq OWNED BY public.passenger.passenger_id;


--
-- Name: driver driver_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.driver ALTER COLUMN driver_id SET DEFAULT nextval('public.driver_driver_id_seq'::regclass);


--
-- Name: order order_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."order" ALTER COLUMN order_id SET DEFAULT nextval('public.order_order_id_seq'::regclass);


--
-- Name: passenger passenger_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.passenger ALTER COLUMN passenger_id SET DEFAULT nextval('public.passenger_passenger_id_seq'::regclass);


--
-- Data for Name: driver; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.driver (driver_id, name, car_type, car_number, score, available, current_location) FROM stdin;
2	Антон	Lada	EB 345 TH 123	4.70	t	{"ing": 22.3061, "lat": 12.9343}
1	Владислав	Volkswagen Polo	A 123 BC 150	4.80	t	{"ing": 30.3061, "lat": 59.9343}
\.


--
-- Data for Name: order; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."order" (order_id, created_at, updated_at, canceled_at, passenger_id, order_status, driver_assigned, arrived_code, address_from, address_to, tariff, selected_services, comment, price, completed_at) FROM stdin;
5	2025-11-26 12:53:14.725263	2025-11-26 12:53:14.725263	2025-11-26 13:02:04.434198	1	pending	\N		ул. Льва Толстого, 16	Московский вокзал	comfort	{pet,luggage_rack}	Встретить у главного входа	500	\N
14	2025-11-28 11:32:16.876228	2025-11-28 11:32:16.876228	2025-11-28 11:33:28.682983	1	canceled	\N		ул. Льва Толстого, 16	Московский вокзал		{pet,luggage_rack,english}	Встретить у главного входа	550	\N
6	2025-11-26 12:58:37.824816	2025-11-26 12:58:37.824816	\N	1	driver_assigned	1		ул. Льва Толстого, 16	Московский вокзал	comfort	{pet,luggage_rack}	Встретить у главного входа	400	\N
13	2025-11-26 16:56:12.40643	2025-11-26 16:56:12.40643	\N	1	driver_assigned	1		ул. Льва Толстого, 16	Московский вокзал	comfort	{pet,luggage_rack,english}	Встретить у главного входа	600	\N
7	2025-11-26 16:44:24.340383	2025-11-26 16:44:24.340383	2025-11-26 16:44:58.37605	1	searching	\N		ул. Льва Толстого, 16	Московский вокзал	comfort	{pet,luggage_rack}	Встретить у главного входа	400	\N
8	2025-11-26 16:51:04.733384	2025-11-26 16:51:04.733384	2025-11-26 16:51:18.974996	1	searching	\N		ул. Льва Толстого, 16	Московский вокзал	comfort	{pet,luggage_rack}	Встретить у главного входа	500	\N
11	2025-11-26 16:55:58.899031	2025-11-26 16:55:58.899031	\N	1	completed	2	0213	ул. Льва Толстого, 16	Московский вокзал	comfort	{pet,luggage_rack,english}	Встретить у главного входа	700	2025-11-28 11:38:49.393028
19	2025-11-30 14:20:45.159745	2025-11-30 14:21:58.833961	2025-11-30 14:21:58.831107	1	canceled	\N		ул. Братьев Кашириных	ТЦ МореМолл	comfort_plus	{pet}	Буду на парковке	350	\N
9	2025-11-26 16:52:26.804584	2025-11-26 16:52:26.804584	2025-11-26 16:52:48.44006	1	canceled	\N		ул. Льва Толстого, 16	Московский вокзал	comfort	{pet,luggage_rack,english}	Встретить у главного входа	700	\N
10	2025-11-26 16:55:37.353918	2025-11-26 16:55:37.353918	2025-11-26 16:58:01.491641	1	canceled	\N		ул. Льва Толстого, 16	Московский вокзал	comfort	{pet,luggage_rack,english}	Встретить у главного входа	600	\N
15	2025-11-28 11:59:44.855595	2025-11-28 11:59:44.855595	\N	1	pending	\N		ул. Льва Толстого, 16	Московский вокзал	comfort	{pet,luggage_rack,english}	Встретить у главного входа	700	\N
16	2025-11-28 12:44:19.184001	2025-11-28 12:44:19.184001	\N	1	pending	\N		ул. Льва Толстого, 16	Московский вокзал	comfort_plus	{pet,luggage_rack,english}	Встретить у главного входа	650	\N
21	2025-12-01 12:06:49.820549	2025-12-01 12:06:49.820549	\N	1	pending	\N		ул. Братьев Кашириных	ТЦ МореМолл	comfort_plus	{pet}	Буду на парковке	350	\N
22	2025-12-01 12:07:26.608504	2025-12-01 12:07:26.608504	\N	1	pending	\N		ул. Братьев Кашириных	ТЦ МореМолл	comfort_plus	{pet}	Буду на парковке	350	\N
23	2025-12-01 12:21:25.51237	2025-12-01 12:21:25.51237	\N	1	pending	\N		ул. Братьев Кашириных	ТЦ МореМолл	comfort_plus	{pet}	Буду на парковке	350	\N
24	2025-12-01 12:25:02.590456	2025-12-01 12:25:02.590456	\N	1	pending	\N		ул. Братьев Кашириных	ТЦ МореМолл	comfort_plus	{pet}	Буду на парковке	350	\N
12	2025-11-26 16:56:06.578211	2025-11-26 16:56:06.578211	\N	1	in_progress	\N		ул. Льва Толстого, 16	Московский вокзал	comfort	{pet,luggage_rack,english}	Встретить у главного входа	700	2025-11-28 11:38:57.263882
26	2025-12-01 12:27:53.581521	2025-12-01 12:27:53.621812	2025-12-01 12:27:53.619721	1	canceled	\N		ул. Братьев Кашириных	ТЦ МореМолл	comfort_plus	{pet}	Буду на парковке	350	\N
17	2025-11-28 14:41:28.126666	2025-11-28 14:41:28.126666	\N	1	completed	2	1302	ул. Льва Толстого, 16	Московский вокзал	comfort_plus	{pet,luggage_rack,english}	Встретить у главного входа	650	2025-11-28 14:46:46.762759
18	2025-11-28 16:00:52.445252	2025-11-28 16:03:06.950502	\N	1	completed	2	2103	ул. Братьев Кашириных	ТЦ МореМолл	comfort_plus	{pet}	Буду на парковке	350	2025-11-28 16:03:06.950502
25	2025-12-01 12:25:42.496152	2025-12-01 12:25:42.520964	2025-12-01 12:25:42.51733	1	canceled	\N		ул. Братьев Кашириных	ТЦ МореМолл	comfort_plus	{pet}	Буду на парковке	350	\N
28	2025-12-01 12:30:01.67247	2025-12-01 12:30:01.716071	2025-12-01 12:30:01.713416	1	canceled	\N		ул. Братьев Кашириных	ТЦ МореМолл	comfort_plus	{pet}	Буду на парковке	350	\N
29	2025-12-01 12:37:25.046294	2025-12-01 12:37:25.046294	\N	1	pending	\N		ул. Ленина 1	ул. Пушкина 10	comfort	{}		350	\N
30	2025-12-01 12:42:12.599349	2025-12-01 12:42:12.599349	\N	1	pending	\N		ул. Ленина 1	ул. Пушкина 10	comfort	{}		350	\N
27	2025-12-01 12:28:24.951441	2025-12-01 12:28:24.972353	2025-12-01 12:28:24.970653	1	canceled	\N		ул. Братьев Кашириных	ТЦ МореМолл	comfort_plus	{pet}	Буду на парковке	350	\N
31	2025-12-01 12:45:21.498	2025-12-01 12:45:21.536616	2025-12-01 12:45:21.53558	1	canceled	2	0231	ул. Ленина 1	ул. Пушкина 10	comfort	{}		250	\N
33	2025-12-01 12:58:02.259627	2025-12-01 12:58:02.277211	\N	1	searching	\N		ул. Ленина 1	ул. Пушкина 10	comfort	{}		250	\N
32	2025-12-01 12:55:57.777862	2025-12-01 12:55:57.798957	\N	1	searching	\N		ул. Ленина 1	ул. Пушкина 10	comfort	{}		350	\N
34	2025-12-01 12:58:17.170993	2025-12-01 12:58:17.18812	\N	1	searching	\N		ул. Ленина 1	ул. Пушкина 10	comfort	{}		350	\N
35	2025-12-01 13:12:24.966112	2025-12-01 13:12:25.004091	\N	1	searching	\N		ул. Ленина 1	ул. Пушкина 10	comfort	{}		350	\N
36	2025-12-01 13:18:58.56484	2025-12-01 14:27:03.74865	\N	1	waiting_for_confirmation	1	1032	ул. Ленина 1	ул. Пушкина 10	comfort	{}		150	\N
38	2025-12-01 13:21:29.050452	2025-12-01 13:21:29.050452	\N	1	pending	\N		ул. Ленина 1	ул. Пушкина 10	comfort	{}		350	\N
37	2025-12-01 13:21:01.835346	2025-12-01 13:21:01.864125	\N	1	searching	\N		ул. Ленина 1	ул. Пушкина 10	comfort	{}		350	\N
39	2025-12-01 13:21:58.268735	2025-12-01 13:21:58.285735	\N	1	searching	\N		ул. Ленина 1	ул. Пушкина 10	comfort	{}		250	\N
40	2025-12-01 13:22:20.762317	2025-12-01 13:22:20.778597	\N	1	searching	\N		ул. Ленина 1	ул. Пушкина 10	comfort	{}		250	\N
42	2025-12-01 13:41:50.287881	2025-12-01 13:41:50.319623	\N	1	searching	\N		ул. Ленина 1	ул. Пушкина 10	comfort	{}		350	\N
43	2025-12-01 13:54:34.318287	2025-12-01 13:54:34.370999	\N	1	waiting_for_confirmation	1	2103	ул. Ленина 1	ул. Пушкина 10	comfort	{}		250	\N
44	2025-12-01 14:01:34.959038	2025-12-01 14:01:35.063534	\N	1	in_progress	1	1032	ул. Ленина 1	ул. Пушкина 10	comfort	{}		250	\N
45	2025-12-01 14:03:59.188699	2025-12-01 14:03:59.237883	\N	1	completed	1	2130	ул. Ленина 1	ул. Пушкина 10	comfort	{}		150	2025-12-01 14:03:59.237883
46	2025-12-01 14:21:34.486667	2025-12-01 14:21:34.486667	\N	1	pending	\N		ул. Братьев Кашириных	ТЦ МореМолл	comfort_plus	{pet}	Буду на парковке	350	\N
\.


--
-- Data for Name: passenger; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.passenger (passenger_id, name, city, phone_number) FROM stdin;
1	Алена	Москва	+79995818012
\.


--
-- Name: driver_driver_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.driver_driver_id_seq', 2, true);


--
-- Name: order_order_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.order_order_id_seq', 46, true);


--
-- Name: passenger_passenger_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.passenger_passenger_id_seq', 1, true);


--
-- Name: driver driver_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.driver
    ADD CONSTRAINT driver_pkey PRIMARY KEY (driver_id);


--
-- Name: order order_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."order"
    ADD CONSTRAINT order_pkey PRIMARY KEY (order_id);


--
-- Name: passenger passenger_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.passenger
    ADD CONSTRAINT passenger_pkey PRIMARY KEY (passenger_id);


--
-- Name: order order_driver_assigned_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."order"
    ADD CONSTRAINT order_driver_assigned_fkey FOREIGN KEY (driver_assigned) REFERENCES public.driver(driver_id);


--
-- Name: order order_passenger_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."order"
    ADD CONSTRAINT order_passenger_id_fkey FOREIGN KEY (passenger_id) REFERENCES public.passenger(passenger_id);


--
-- PostgreSQL database dump complete
--

