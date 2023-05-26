export interface Response<TData = any> {
  success: boolean;
  data: TData;
}

export interface Data {
  now_playing: Music;
  list: Music[];
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
