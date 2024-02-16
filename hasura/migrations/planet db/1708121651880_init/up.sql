SET check_function_bodies = false;
CREATE FUNCTION public.set_current_timestamp_updated_at() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
DECLARE
  _new record;
BEGIN
  _new := NEW;
  _new."updated_at" = NOW();
  RETURN _new;
END;
$$;
CREATE TABLE public.moons (
    id integer NOT NULL,
    name text NOT NULL,
    planet_id integer NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
CREATE SEQUENCE public.moons_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.moons_id_seq OWNED BY public.moons.id;
CREATE TABLE public.planets (
    id integer NOT NULL,
    name text NOT NULL,
    rotation bigint NOT NULL,
    revolution bigint NOT NULL,
    radius bigint NOT NULL,
    temperature integer NOT NULL,
    overview_content text NOT NULL,
    overview_source text NOT NULL,
    structure_content text NOT NULL,
    structure_source text NOT NULL,
    geology_content text NOT NULL,
    geology_source text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
CREATE SEQUENCE public.planets_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.planets_id_seq OWNED BY public.planets.id;
ALTER TABLE ONLY public.moons ALTER COLUMN id SET DEFAULT nextval('public.moons_id_seq'::regclass);
ALTER TABLE ONLY public.planets ALTER COLUMN id SET DEFAULT nextval('public.planets_id_seq'::regclass);
ALTER TABLE ONLY public.moons
    ADD CONSTRAINT moons_name_key UNIQUE (name);
ALTER TABLE ONLY public.moons
    ADD CONSTRAINT moons_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.planets
    ADD CONSTRAINT planets_name_key UNIQUE (name);
ALTER TABLE ONLY public.planets
    ADD CONSTRAINT planets_pkey PRIMARY KEY (id);
CREATE TRIGGER set_public_moons_updated_at BEFORE UPDATE ON public.moons FOR EACH ROW EXECUTE FUNCTION public.set_current_timestamp_updated_at();
COMMENT ON TRIGGER set_public_moons_updated_at ON public.moons IS 'trigger to set value of column "updated_at" to current timestamp on row update';
CREATE TRIGGER set_public_planets_updated_at BEFORE UPDATE ON public.planets FOR EACH ROW EXECUTE FUNCTION public.set_current_timestamp_updated_at();
COMMENT ON TRIGGER set_public_planets_updated_at ON public.planets IS 'trigger to set value of column "updated_at" to current timestamp on row update';
ALTER TABLE ONLY public.moons
    ADD CONSTRAINT moons_planet_id_fkey FOREIGN KEY (planet_id) REFERENCES public.planets(id);
