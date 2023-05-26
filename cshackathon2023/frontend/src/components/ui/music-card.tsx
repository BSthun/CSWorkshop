import { Music } from "@/types/api";
import dayjs from "@/utils/time";
import { Box, IconButton, Stack, Typography } from "@mui/material";
import React, { useCallback } from "react";
import { BiMinusCircle } from "react-icons/bi";
import { FaPlay } from "react-icons/fa";
import { MdAccountCircle } from "react-icons/md";

export interface MusicCardProps extends Partial<Music> {
  heading?: boolean;
}

export default function MusicCard({
  heading = false,
  artist = "Taylor Swift",
  artwork_url = `https://source.unsplash.com/random/400×400/?music`,
  is_owned = false,
  is_playing = false,
  queue_at = new Date().toISOString(),
  queue_by = "Sirawit",
  title = "This is a long string that is OK to truncate please and thank youasdasdasdasd",
}: MusicCardProps) {
  const [size, setSize] = React.useState(heading ? 70 : 56);
  const [loaded, setLoaded] = React.useState(false);
  const onLoaded = useCallback(() => {
    setLoaded(true);
  }, []);
  const sizeRef = React.useRef<HTMLDivElement>(null);

  React.useLayoutEffect(() => {
    if (sizeRef.current) {
      const size = sizeRef.current.clientHeight;
      setSize(size + (heading ? 36 : 4));
    }
  }, [sizeRef.current?.clientHeight, heading, is_playing]);

  const timeFromNow = dayjs(queue_at).toNow();

  return (
    <Box
      display="flex"
      alignItems="center"
      gap={2}
      py={heading ? 1.5 : 2}
      pl={3}
      pr={5}
      position="relative"
      // sx={{ background: is_playing ? "#F2F2F2" : "transparent" }}
    >
      <Box
        m={1}
        sx={{
          width: `${size}px`,
          height: `${size}px`,
          borderRadius: "12px",
          background: "gray",
          overflow: "hidden",
          transition: "all .5s ease-in-out",
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
          opacity: loaded ? 1 : 0,
          position: "relative",
        }}
      >
        <Box
          sx={{
            transition: "all .5s ease-in-out",
            position: "absolute",
            backdropFilter: "blur(4px)",
            opacity: is_playing ? 1 : 0,
            inset: 0,
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
            fontSize: "24px",
            color: "white",
          }}
        >
          {is_playing && <FaPlay className="animated-pop-in" />}
        </Box>
        <img
          src={artwork_url}
          alt="artwork_url"
          style={{ objectFit: "contain", height: "100%" }}
          onLoad={onLoaded}
        />
      </Box>

      <Stack
        display="flex"
        flex={1}
        minWidth="0px"
        gap={heading ? 1 : 0.5}
        justifyContent="center"
        ref={sizeRef}
      >
        <Typography
          className="text-overflow"
          fontSize={heading ? "17px" : "16px"}
          lineHeight={1}
          color="#050505"
          fontWeight="800"
        >
          {title}
        </Typography>
        <Typography
          className="text-overflow"
          fontSize={heading ? "16px" : "14px"}
          lineHeight={1}
          color="#050505"
          fontWeight="600"
        >
          {artist}
        </Typography>
        {!heading && (
          <Typography
            className="text-overflow"
            fontSize="14px"
            lineHeight={1}
            color="#919191"
            fontWeight="600"
            mt={0.25}
            sx={{
              display: "flex",
              alignItems: "center",
              gap: 0.5,
            }}
          >
            <MdAccountCircle /> {is_owned ? "You" : queue_by} • {timeFromNow}
          </Typography>
        )}
      </Stack>
      {is_owned && (
        <IconButton size="small" sx={{ position: "absolute", right: 5 }}>
          <BiMinusCircle />
        </IconButton>
      )}
    </Box>
  );
}
