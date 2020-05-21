/*
 * Copyright DataStax, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package stargate.service.config

import scala.beans.BeanProperty
import com.datastax.oss.driver.api.core.auth.AuthProvider
import com.datastax.oss.driver.internal.core.auth.PlainTextAuthProvider
import com.datastax.oss.driver.api.core.context.DriverContext

object ParsedStargateConfig{
  var globalConfig: ParsedStargateConfig = _
}
final case class ParsedStargateConfig(
                                 @BeanProperty val httpPort: Int,
                                 @BeanProperty val defaultTTL: Int,
                                 @BeanProperty val defaultLimit: Int,
                                 @BeanProperty val maxSchemaSizeKB: Long,
                                 @BeanProperty val maxRequestSizeKB: Long,
                                 @BeanProperty val maxMutationSizeKB: Long,
                                 @BeanProperty val cassandraContactPoints: List[(String, Int)],
                                 @BeanProperty val cassandraDataCenter: String,
                                 @BeanProperty val cassandraReplication: Int,
                                 @BeanProperty val stargateKeyspace: String,
                                 @BeanProperty val cassandraAuthProvider: String,
                                 @BeanProperty val cassandraUserName: String,
                                 @BeanProperty val cassandraPassword: String
)