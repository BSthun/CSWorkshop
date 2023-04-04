import {
	Container,
	Grid,
	Paper,
	Stack,
} from '@mui/material';
import React from 'react';

const arr = [0, 0, 0, 0, 0, 0, 0, 0,0];
const Youtube = () => {
	return (
		<Stack>
			<Paper square sx={{ height: 64, zIndex: 20 }} elevation={2}>
				<iframe src="https://csc105-workshop.bsthun.com/components/comp18"
				        height="100%"
				        width="100%"
				        frameBorder="0"
				/>
			</Paper>
			<Stack flexDirection="row" height="calc(100vh - 64px)">
				<Paper square sx={{ width: 300 }} elevation={2}>
					<iframe src="https://csc105-workshop.bsthun.com/components/comp19"
					        height="100%"
					        width="100%"
					        frameBorder="0"
					/>
				</Paper>
				<Container maxWidth="lg">
					<Grid container>
						{
							arr.map((el) => (
								<Grid item sm={4}>
									<iframe src="https://csc105-workshop.bsthun.com/components/comp20"
									        height="300px"
									        width="100%"
									        frameBorder="0"
									/>
								</Grid>
							))
						}
					</Grid>
				</Container>
			</Stack>
		</Stack>
	);
};

export default Youtube;
