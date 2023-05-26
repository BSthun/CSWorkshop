import MusicCard from "@/components/ui/music-card";
import { LoadingButton } from "@mui/lab";
import { Box, Divider, Stack, Typography } from "@mui/material";
import { BiPlusCircle } from "react-icons/bi";
import { FaPlay } from "react-icons/fa";

const boxShadow = `
0 1px 1px hsl(0deg 0% 0% / 0.025),
0 2px 2px hsl(0deg 0% 0% / 0.025),
0 4px 4px hsl(0deg 0% 0% / 0.025),
0 8px 8px hsl(0deg 0% 0% / 0.025),
0 16px 16px hsl(0deg 0% 0% / 0.025)
`;

export default function MainView() {
  return (
    <Stack sx={{ height: "100vh", display: "flex" }}>
      <Stack
        sx={{
          position: "fixed",
          top: 0,
          left: 0,
          right: 0,
          background: "#F2F2F2C8",
          backdropFilter: "blur(8px)",
          borderRadius: "0px 0px 10px 10px",
          boxShadow,
          border: "1px solid rgba(33,33,33,.1)",
          zIndex: 5,
          height: "12rem",
        }}
      >
        <Typography fontSize="24px" fontWeight="500" py={2} align="center">
          CS Hackathon 2023
        </Typography>
        <MusicCard heading />
      </Stack>
      <Box flex={1} overflow="auto" pb="5.5rem" pt={"12rem"}>
        <Stack divider={<Divider />}>
          <MusicCard is_playing />
          <MusicCard />
          <MusicCard />
          <MusicCard is_owned />
          <MusicCard />
          <MusicCard />
          <MusicCard is_owned />
          <MusicCard />
          <MusicCard />
          <MusicCard is_owned />
          <MusicCard />
          <MusicCard />
          <MusicCard is_owned />
          <MusicCard />
          <MusicCard />
          <MusicCard is_owned />
          <MusicCard />
          <MusicCard />
          <MusicCard is_owned />
          <MusicCard />
          <MusicCard />
          <MusicCard is_owned />
          <MusicCard />
          <MusicCard />
          <MusicCard is_owned />
          <MusicCard />
        </Stack>
      </Box>
      <Box
        sx={{
          position: "fixed",
          left: 0,
          bottom: 0,
          right: 0,
          background: "#F2F2F2C8",
          backdropFilter: "blur(8px)",
          borderRadius: "10px 10px 0px 0px",
          boxShadow,
          border: "1px solid rgba(33,33,33,.1)",
          display: "flex",
          paddingLeft: 2,
          paddingRight: 2,
          gap: 2,
          height: "5.5rem",
          alignItems: "center",
          zIndex: 5,
        }}
      >
        <LoadingButton
          variant="contained"
          color="success"
          sx={{
            borderRadius: "12px",
            width: "55px",
            height: "55px",
            fontSize: "36px",
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
          }}
          disableElevation
        >
          <BiPlusCircle />
        </LoadingButton>
        <LoadingButton
          variant="contained"
          color="primary"
          sx={{
            borderRadius: "12px",
            width: "55px",
            height: "55px",
            fontSize: "24px",
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
          }}
          disableElevation
        >
          <FaPlay />
        </LoadingButton>
        <Stack flex={1} justifyContent="center">
          <Typography fontSize="14px">32/46 attendees voted play</Typography>
          <Typography fontSize="14px">You’re currently voted pause</Typography>
        </Stack>
      </Box>
    </Stack>
  );
}
