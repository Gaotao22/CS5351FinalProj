package devils.scrumtool.repositories;

import devils.scrumtool.entities.PBI;
// Java libraries
import java.util.List;
// Spring libraries
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.CrudRepository;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;
import org.springframework.transaction.annotation.Transactional;

@Repository
public interface PBIRepository extends CrudRepository<PBI, Integer> {
    // Query Methods
    List<PBI> findByProject_IdAndIsEpic(Integer projectId, Boolean isEpic);

    List<PBI> findByProject_IdAndEpicId(Integer projectId, Integer epicId);

    List<PBI> findBySprint_Id(Integer sprintId);

    PBI getSprint_IdById(Integer pbiId);

    @Modifying
    @Query(value = "UPDATE pbis AS p SET totaltime = (SELECT SUM(comtime) FROM tasks AS t WHERE t.pbi_id=p.id) WHERE p.id = ?1", nativeQuery = true)
    // @Query("UPDATE PBI p SET p.totaltime = :pbiId WHERE p.id = :pbiId")
    @Transactional
    void updateTotalTime(@Param("pbiId") Integer pbiId);

    @Transactional
    void deleteById(Integer pbiId);
}
