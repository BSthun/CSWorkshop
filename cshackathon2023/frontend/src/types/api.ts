export interface Response<TData = any> {
  success: boolean;
  data: TData;
  message?: string;
}

export interface Data {
  now_playing: Music;
  items: Music[];
  control: Control;
}

export interface Control {
  voted_play: boolean;
  attendee_voted_play: boolean;
  attendee_count: number;
}

export interface Music {
  id: number;
  artwork_url: string;
  title: string;
  artist: string;
  queue_by: string;
  queue_at: string;
  is_playing: boolean;
  is_owned: boolean;
}

export type MusicStateResponse = Response<Data>;

export interface SeachData {
  items: SearchMusic[];
}

export interface SearchMusic {
  album: string;
  artist: string;
  artwork_url: string;
  spotify_id: string;
  title: string;
}

export type SearchMusicResponse = Response<SeachData>;
