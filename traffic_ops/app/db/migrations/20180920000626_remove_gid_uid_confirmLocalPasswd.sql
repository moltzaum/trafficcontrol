/*
	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at
	    http://www.apache.org/licenses/LICENSE-2.0
	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

-- we are going to clean up capabilities for a reseed
-- it is really just a few, but I'm deleting everything just so I don't miss anything..
ALTER TABLE tm_user DROP COLUMN gid;
ALTER TABLE tm_user DROP COLUMN uid;
ALTER TABLE tm_user DROP COLUMN confirm_local_passwd;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE tm_user ADD COLUMN gid bigint;
ALTER TABLE tm_user ADD COLUMN uid bigint;
ALTER TABLE tm_user ADD COLUMN confirm_local_passwd text;

