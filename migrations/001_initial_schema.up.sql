CREATE TYPE gender AS ENUM ('m', 'f');

CREATE TABLE public.user (
  id integer generated always as identity primary key,
  first_name character varying(255),
  last_name character varying(255),
  username character varying(255),
  email character varying(255),
  password text,
  gender gender,
  locale character varying(5),
  avatar_url text,
  active boolean,
  email_verified boolean,
  failed_attempts integer,
  last_login_at timestamp with time zone,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone
);

CREATE TABLE public.role (
  id integer generated always as identity primary key,
  name character varying(30)
);

CREATE TABLE public.user_role (
  user_id integer not null,
  role_id integer not null,
  primary key (user_id, role_id),
  foreign key (user_id) references public.user (id),
  foreign key (role_id) references public.role (id)
);

CREATE TABLE public.address (
  id integer generated always as identity primary key,
  user_id integer,
  country character varying(255),
  city character varying(255),
  address_1 character varying(255),
  address_2 character varying(255),
  zip character varying(30),
  longitude numeric(11, 8),
  latitude numeric(11, 8),
  foreign key (user_id) references public.user (id)
);

CREATE TABLE public.item (
  id integer generated always as identity primary key,
  user_id integer,
  type_id integer,
  size_id integer,
  color_id integer,
  description text,
  sku character varying(30)
);

CREATE TABLE public.item_info (
  id integer generated always as identity primary key,
  item_id integer,
  price integer,
  description text,
  foreign key (item_id) references public.item (id)
);

CREATE TABLE public.item_images (
  id integer generated always as identity primary key,
  item_id integer,
  url text,  
  foreign key (item_id) references public.item (id)
);

CREATE TABLE public.related_item (
  id integer generated always as identity primary key,
  item_id integer,
  related_item_id integer,  
  foreign key (item_id) references public.item (id),
  foreign key (related_item_id) references public.item (id)
);

CREATE TABLE public.manufacturer (
  id integer generated always as identity primary key,
  name character varying(100),
  type character varying(50),
  email text,
  website_url text,
  address text,
  description text
);

CREATE TABLE public.token (
  id integer generated always as identity primary key,
  user_id integer,
  reset_token text,
  reset_expires timestamp with time zone,
  verification_token text,
  verification_expires timestamp with time zone,
  refresh_token text,
  refresh_expires timestamp with time zone,
  refresh_revoked boolean,
  foreign key (user_id) references public.user (id)
);

CREATE TABLE public.color (
  id integer generated always as identity primary key
);

CREATE TABLE public.size (
  id integer generated always as identity primary key
);
