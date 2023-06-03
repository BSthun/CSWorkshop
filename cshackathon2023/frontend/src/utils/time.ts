import Dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";

Dayjs.extend(relativeTime);

const dayjs = Dayjs;

export default dayjs;
