import { Box, IconButton, SwipeableDrawer, TextField } from "@mui/material";
import React from "react";
import MusicCard from "./music-card";
import { MdClose } from "react-icons/md";
import { SearchMusic } from "@/types/api";
import axios from "axios";

type SearchDrawerProps = {
  open: boolean;
  setOpenDrawer: React.Dispatch<React.SetStateAction<boolean>>;
  searchInput: string;
  setSearchInput: React.Dispatch<React.SetStateAction<string>>;
  songList: SearchMusic[];
};

const SearchDrawer: React.FC<SearchDrawerProps> = ({
  open,
  setOpenDrawer,
  searchInput,
  setSearchInput,
  songList,
}) => {
  const queue = (spotify_id: any) => {
    axios
      .post("/api/music/queue", {
        trackId: spotify_id,
      })
      .then((response) => {
        if (response.data.success) {
          alert("Song queued!");
          setOpenDrawer(false);
        } else {
          alert("Error: " + response.data.message);
        }
      });
  };
  return (
    <SwipeableDrawer
      anchor="bottom"
      open={open}
      onClose={() => {
        setOpenDrawer(false);
        setSearchInput("");
      }}
      onOpen={() => setOpenDrawer(true)}
    >
      <Box
        sx={{
          width: "100%",
          backgroundColor: "#FAFAFA",
          height: "70vh",
          pb: 2,
          px: 2,
        }}
      >
        <Box
          sx={{
            width: "100%",
            position: "fixed",
            backgroundColor: "#FAFAFA",
            py: 2,
            pr: 4,
            zIndex: 1,
          }}
        >
          <TextField
            value={searchInput}
            onChange={(e) => setSearchInput(e.target.value)}
            variant="outlined"
            placeholder="Find some musics"
            fullWidth
            InputProps={{
              sx: {
                height: "50px",
                pr: 1,
              },
              endAdornment: searchInput.length > 0 && (
                <IconButton size="small" onClick={() => setSearchInput("")}>
                  <MdClose />
                </IconButton>
              ),
            }}
          />
        </Box>

        <Box mt={"80px"} zIndex={-1}>
          {songList.map((song) => (
            <MusicCard search key={song.spotify_id} {...song} queue={queue} />
          ))}
        </Box>
      </Box>
    </SwipeableDrawer>
  );
};

export default SearchDrawer;
